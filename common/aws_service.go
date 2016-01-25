package common

import (
	"../config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elb"
)

func ElbService() *elb.ELB {
	conf := config.GetConfig()

	return elb.New(
		session.New(),
		&aws.Config{
			Region:      aws.String(conf.Aws.Region),
			Credentials: conf.AwsCredentials(),
		},
	)
}

func Ec2Service() *ec2.EC2 {
	conf := config.GetConfig()

	return ec2.New(
		session.New(),
		&aws.Config{
			Region:      aws.String(conf.Aws.Region),
			Credentials: conf.AwsCredentials(),
		},
	)
}

func findVpcs(params *ec2.DescribeVpcsInput) ([]*ec2.Vpc, error) {
	service := Ec2Service()

	res, err := service.DescribeVpcs(params)
	if err != nil {
		// With vpcid request, aws-sdk return error with 'InvalidVpcID.NotFound'.
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "InvalidVpcID.NotFound" {
				return []*ec2.Vpc{}, &VpcNotFoundError{}
			}
		}

		return []*ec2.Vpc{}, err
	}

	if len(res.Vpcs) == 0 {
		return []*ec2.Vpc{}, &VpcNotFoundError{}
	}

	return res.Vpcs, nil
}

func ShowError(err error) {
	if awsErr, ok := err.(awserr.Error); ok {
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			logger.Error(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		} else {
			// Generic AWS Error with Code, Message, and original error (if any)
			logger.Error(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		}
	} else {
		// This case should never be hit, the SDK should always return an error
		// which satisfies the awserr.Error interface.
		logger.Error("%s\n", err.Error())
	}
}

func IsNetworkAccessible(instance *ec2.Instance) bool {
	if *instance.State.Name != "running" {
		logger.Warn("Instance(%s) is not running.\n", *instance.InstanceId)
		return false
	}

	if instance.PublicIpAddress == nil {
		logger.Warn("Instance(%s) does not have Public IPAddress.\n", *instance.InstanceId)
		return false
	}

	return true
}
