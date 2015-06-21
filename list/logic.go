package list

import (
	"../config"
	"../formatter"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
)

// Reference Code
// http://qiita.com/draco/items/59c94ab8c66314d3a9f0

// Sample Code
// https://github.com/aws/aws-sdk-go/blob/7e9078c250876f26da48aaf36b8dce6a462ecd8a/service/ec2/examples_test.go#L2876

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

func describeParams() *ec2.DescribeInstancesInput {
	return &ec2.DescribeInstancesInput{}
}

func ec2Service() *ec2.EC2 {
	conf := config.GetConfig()

	return ec2.New(
		&aws.Config{
			Region:      conf.Aws.Region,
			Credentials: conf.AwsCredentials(),
		},
	)
}

func newTableWriter(writer io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(writer)

	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetColumnSeparator("\t")
	table.SetColWidth(80)

	return table
}

func ShowEc2Instances(writer io.Writer) int {
	res, err := ec2Service().DescribeInstances(
		describeParams(),
	)

	if err != nil {
		fmt.Println("failed...")
		showError(err)
		return 1
	}

	table := newTableWriter(writer)
	formatter := formatter.NewEc2Formatter()

	for _, r := range res.Reservations {
		for _, i := range r.Instances {
			table.Append(formatter.Format(i))
		}
	}

	table.Render()
	return 0
}
