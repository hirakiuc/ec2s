package list

import (
	"io"

	"github.com/aws/aws-sdk-go/service/ec2"

	"../../cache"
	"../../common"
	"../../formatter"
)

// Reference Code
// http://qiita.com/draco/items/59c94ab8c66314d3a9f0

// Sample Code
// https://github.com/aws/aws-sdk-go/blob/7e9078c250876f26da48aaf36b8dce6a462ecd8a/service/ec2/examples_test.go#L2876

func vpcId(instance *ec2.Instance) string {
	if instance.VPCID == nil {
		return ""
	} else {
		return *instance.VPCID
	}
}

func loadVpcCache() *cache.VpcCache {
	cache := cache.GetVpcCache()
	if cache.MakeCache() == false {
		return nil
	} else {
		return cache
	}
}

func ShowEc2Instances(writer io.Writer, options common.FilterInterface) int {
	vpcCache := loadVpcCache()
	if vpcCache == nil {
		logger.Error("failed to make vpc cache.\n")
		return 10
	}
	instanceCache := cache.GetEc2InstanceCache()

	service := common.Ec2Service()
	params, err := options.InstancesFilter()
	if err != nil {
		common.ShowError(err)
		return 1
	}

	res, err := service.DescribeInstances(params)
	if err != nil {
		logger.Error("failed to load EC2 instances.\n")
		common.ShowError(err)
		return 1
	}

	table := common.NewTableWriter(writer)
	formatter := formatter.NewEc2Formatter()

	for _, r := range res.Reservations {
		for _, i := range r.Instances {
			vpc := vpcCache.ReadEntry(vpcId(i))
			table.Append(formatter.Format(vpc, i))

			instanceCache.WriteEntry(
				*i.InstanceID,
				i,
			)
		}
	}

	table.Render()
	return 0
}
