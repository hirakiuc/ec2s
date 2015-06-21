package vpcs

import (
	"../config"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
)

func showError(err error) {
	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		// This case should never be hit, the SDK should always return an error
		// which satisfies the awserr.Error interface.
		fmt.Println(err.Error())
	}
}

func describeVpc(writer *tablewriter.Table, v *ec2.VPC) {
	var tag_name string
	for _, t := range v.Tags {
		if *t.Key == "Name" {
			tag_name = *t.Value
		}
	}

	writer.Append([]string{
		tag_name,
		*v.VPCID,
		*v.State,
	})
}

func (c *Command) showVpcs(writer io.Writer) {
	conf := config.GetConfig()

	credentials := credentials.NewStaticCredentials(
		conf.Aws.AccessKeyId,
		conf.Aws.SecretAccessKey,
		"",
	)

	svc := ec2.New(
		&aws.Config{
			Region:      conf.Aws.Region,
			Credentials: credentials,
		},
	)

	res, err := svc.DescribeVPCs(nil)
	if err != nil {
		fmt.Println("failed...")
		showError(err)
		os.Exit(1)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetColumnSeparator("\t")
	table.SetColWidth(80)

	for _, vpc := range res.VPCs {
		describeVpc(table, vpc)
	}

	table.Render()
}
