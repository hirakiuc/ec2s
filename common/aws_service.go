package common

import (
	"../config"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func Ec2Service() *ec2.EC2 {
	conf := config.GetConfig()

	return ec2.New(
		&aws.Config{
			Region:      conf.Aws.Region,
			Credentials: conf.AwsCredentials(),
		},
	)
}

func vpcByName(vpcName string) (*ec2.VPC, error) {
	service := Ec2Service()

	params := &ec2.DescribeVPCsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(vpcName),
				},
			},
		},
	}

	res, err := service.DescribeVPCs(params)
	if err != nil {
		fmt.Println("failed to fetch vpc.")
		ShowError(err)
		return nil, err
	}

	if len(res.VPCs) > 0 {
		return res.VPCs[0], nil
	} else {
		fmt.Printf("No such vpc found: %s", vpcName)
		return nil, nil
	}
}

func ShowError(err error) {
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
