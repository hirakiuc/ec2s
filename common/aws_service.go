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

func findVpcs(params *ec2.DescribeVPCsInput) ([]*ec2.VPC, error) {
	service := Ec2Service()

	res, err := service.DescribeVPCs(params)
	if err != nil {
		// With vpcid request, aws-sdk return error with 'InvalidVpcID.NotFound'.
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "InvalidVpcID.NotFound" {
				return []*ec2.VPC{}, &VpcNotFoundError{}
			}
		}

		return []*ec2.VPC{}, err
	}

	if len(res.VPCs) == 0 {
		return []*ec2.VPC{}, &VpcNotFoundError{}
	}

	return res.VPCs, nil
}

func ShowError(err error) {
	if awsErr, ok := err.(awserr.Error); ok {
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		} else {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		}
	} else {
		// This case should never be hit, the SDK should always return an error
		// which satisfies the awserr.Error interface.
		fmt.Println(err.Error())
	}
}

func IsNetworkAccessible(instance *ec2.Instance) bool {
	if *instance.State.Name != "running" {
		fmt.Printf("The instance is not running. (%s)\n", *instance.InstanceID)
		return false
	}

	if instance.PublicIPAddress == nil {
		fmt.Printf("The instance does not have Public IPAddress. (%s)\n", *instance.InstanceID)
		return false
	}

	return true
}
