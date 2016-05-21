package elbs

import (
	"io"

	"github.com/aws/aws-sdk-go/service/elb"

	"github.com/hirakiuc/ec2s/internal/cache"
	"github.com/hirakiuc/ec2s/internal/common"
	"github.com/hirakiuc/ec2s/internal/formatter"
)

func loadVpcCache() error {
	cache := cache.GetVpcCache()
	if err := cache.MakeCache(); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Command) showElbs(writer io.Writer) error {
	if err := loadVpcCache(); err != nil {
		return err
	}

	service := common.ElbService()

	params := &elb.DescribeLoadBalancersInput{}
	res, err := service.DescribeLoadBalancers(params)
	if err != nil {
		logger.Error("failed to fetch elbs.\n")
		return err
	}

	table := common.NewTableWriter(writer)
	formatter := formatter.NewElbFormatter()

	for _, elb := range res.LoadBalancerDescriptions {
		table.Append(formatter.Format(elb))
	}

	table.Render()
	return nil
}
