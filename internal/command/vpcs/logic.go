package vpcs

import (
	"io"

	"github.com/hirakiuc/ec2s/internal/common"
	"github.com/hirakiuc/ec2s/internal/formatter"
)

func (c *Command) showVpcs(writer io.Writer) error {
	service, err := common.Ec2Service()
	if err != nil {
		return err
	}

	res, err := service.DescribeVpcs(nil)
	if err != nil {
		logger.Error("failed to fetch vpcs.\n")

		return err
	}

	table := common.NewTableWriter(writer)
	formatter := formatter.NewVpcFormatter()

	for _, vpc := range res.Vpcs {
		table.Append(formatter.Format(vpc))
	}

	table.Render()

	return nil
}
