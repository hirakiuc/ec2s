package list

import (
	"io"

	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/hirakiuc/ec2s/internal/cache"
	"github.com/hirakiuc/ec2s/internal/common"
	"github.com/hirakiuc/ec2s/internal/formatter"
)

// Reference Code
// http://qiita.com/draco/items/59c94ab8c66314d3a9f0

// Sample Code
// https://github.com/aws/aws-sdk-go/blob/7e9078c250876f26da48aaf36b8dce6a462ecd8a/service/ec2/examples_test.go#L2876

func vpcID(instance *ec2.Instance) string {
	if instance.VpcId == nil {
		return ""
	}

	return *instance.VpcId
}

func loadVpcCache() (*cache.VpcCache, error) {
	cache := cache.GetVpcCache()
	if err := cache.MakeCache(); err != nil {
		return nil, err
	}

	return cache, nil
}

// ShowEc2Instances shows EC2 instances.
func ShowEc2Instances(writer io.Writer, options common.FilterInterface) error {
	vpcCache, err := loadVpcCache()
	if err != nil {
		return err
	}

	instanceCache := cache.GetEc2InstanceCache()

	service, err := common.Ec2Service()
	if err != nil {
		return err
	}

	params, err := options.InstancesFilter()
	if err != nil {
		return err
	}

	res, err := service.DescribeInstances(params)
	if err != nil {
		logger.Error("failed to load EC2 instances.\n")

		return err
	}

	table := common.NewTableWriter(writer)
	formatter := formatter.NewEc2Formatter()

	for _, r := range res.Reservations {
		for _, i := range r.Instances {
			vpc := vpcCache.ReadEntry(vpcID(i))
			table.Append(formatter.Format(vpc, i))

			instanceCache.WriteEntry(
				*i.InstanceId,
				i,
			)
		}
	}

	table.Render()

	return nil
}
