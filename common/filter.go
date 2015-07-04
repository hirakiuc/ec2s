package common

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type InstanceFilter struct {
	VpcId   string
	VpcName string
}

type FilterInterface interface {
	VpcFilter() (*ec2.Filter, error)
	InstancesFilter() (*ec2.DescribeInstancesInput, error)
}

func (filter *InstanceFilter) vpcIdForFilter() (*string, error) {
	if len(filter.VpcId) > 0 {
		return &filter.VpcId, nil
	}

	if len(filter.VpcName) > 0 {
		vpc, err := vpcByName(filter.VpcName)
		if err != nil {
			return nil, err
		} else {
			return vpc.VPCID, err
		}
	}

	return nil, nil
}

func (filter *InstanceFilter) VpcFilter() (*ec2.Filter, error) {
	vpcId, err := filter.vpcIdForFilter()
	if err != nil {
		return nil, err
	}

	if vpcId == nil {
		return nil, nil
	}

	return &ec2.Filter{
		Name: aws.String("vpc-id"),
		Values: []*string{
			aws.String(*vpcId),
		},
	}, nil
}

func (filter *InstanceFilter) InstancesFilter() (*ec2.DescribeInstancesInput, error) {
	//func InstancesFilter(options FilterInterface) *ec2.DescribeInstancesInput {
	filters := []*ec2.Filter{}

	vpcFilter, err := filter.VpcFilter()
	if err != nil {
		return nil, err
	}

	if vpcFilter != nil {
		filters = append(filters, vpcFilter)
	}

	// If Filters was empty array, aws-sdk-go return error on ec2.DescribeInstances method.
	// So this code avoid the error.
	if len(filters) > 0 {
		return &ec2.DescribeInstancesInput{
			Filters: filters,
		}, nil
	} else {
		return &ec2.DescribeInstancesInput{}, nil
	}
}
