package elbs

import (
	"io"

	"github.com/aws/aws-sdk-go/service/elb"

	"../../common"
	"../../formatter"
)

func (c *Command) showElbs(writer io.Writer) error {
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
