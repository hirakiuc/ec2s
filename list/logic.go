package list

import (
	"../config"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
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

func describeInstance(writer *tablewriter.Table, i *ec2.Instance) {
	var tag_name string
	for _, t := range i.Tags {
		if *t.Key == "Name" {
			tag_name = *t.Value
		}
	}

	var ipAddress string
	if i.PublicIPAddress == nil {
		ipAddress = "---"
	} else {
		ipAddress = *i.PublicIPAddress
	}

	writer.Append([]string{
		tag_name,
		*i.InstanceID,
		*i.InstanceType,
		*i.Placement.AvailabilityZone,
		ipAddress,
		*i.State.Name,
	})
}

func ShowEc2Instances(writer io.Writer) int {
	conf := config.GetConfig()

	credentials := credentials.NewStaticCredentials(
		conf.Aws.AccessKeyId,
		conf.Aws.SecretAccessKey,
		"")

	svc := ec2.New(
		&aws.Config{
			Region:      "ap-northeast-1",
			Credentials: credentials,
		},
	)

	params := &ec2.DescribeInstancesInput{}

	res, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("failed...")
		showError(err)
		return 1
	}

	table := tablewriter.NewWriter(writer)
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetColumnSeparator("\t")
	table.SetColWidth(80)

	for _, r := range res.Reservations {
		for _, i := range r.Instances {
			describeInstance(table, i)
		}
	}

	table.Render()
	return 0
}
