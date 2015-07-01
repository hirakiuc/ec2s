package common

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type InstanceFilter struct {
	VpcId string
}

type FilterInterface interface {
	VpcFilter() *ec2.Filter
}

func (filter *InstanceFilter) VpcFilter() *ec2.Filter {
	if len(filter.VpcId) == 0 {
		return nil
	}

	return &ec2.Filter{
		Name: aws.String("vpc-id"),
		Values: []*string{
			aws.String(filter.VpcId),
		},
	}
}

func InstancesFilter(options FilterInterface) *ec2.DescribeInstancesInput {
	filters := []*ec2.Filter{}

	vpcFilter := options.VpcFilter()
	if vpcFilter != nil {
		filters = append(filters, vpcFilter)
	}

	// If Filters was empty array, aws-sdk-go return error on ec2.DescribeInstances method.
	// So this code avoid the error.
	if len(filters) > 0 {
		return &ec2.DescribeInstancesInput{
			Filters: filters,
		}
	} else {
		return &ec2.DescribeInstancesInput{}
	}
}
