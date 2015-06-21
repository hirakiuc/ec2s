package list

import (
	"../common"
	"../formatter"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// Reference Code
// http://qiita.com/draco/items/59c94ab8c66314d3a9f0

// Sample Code
// https://github.com/aws/aws-sdk-go/blob/7e9078c250876f26da48aaf36b8dce6a462ecd8a/service/ec2/examples_test.go#L2876

func describeParams() *ec2.DescribeInstancesInput {
	return &ec2.DescribeInstancesInput{}
}

func ShowEc2Instances(writer io.Writer) int {
	service := common.Ec2Service()
	res, err := service.DescribeInstances(describeParams())

	if err != nil {
		fmt.Println("failed...")
		common.ShowError(err)
		return 1
	}

	table := common.NewTableWriter(writer)
	formatter := formatter.NewEc2Formatter()

	for _, r := range res.Reservations {
		for _, i := range r.Instances {
			table.Append(formatter.Format(i))
		}
	}

	table.Render()
	return 0
}
