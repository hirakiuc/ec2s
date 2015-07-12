package vpcs

import (
	"io"

	"../../common"
	"../../formatter"
)

func (c *Command) showVpcs(writer io.Writer) int {
	service := common.Ec2Service()
	res, err := service.DescribeVPCs(nil)
	if err != nil {
		logger.Error("failed to fetch vpcs: %v\n")
		common.ShowError(err)
		return 1
	}

	table := common.NewTableWriter(writer)
	formatter := formatter.NewVpcFormatter()

	for _, vpc := range res.VPCs {
		table.Append(formatter.Format(vpc))
	}

	table.Render()
	return 0
}
