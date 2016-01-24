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
	VpcFilterExist() bool
}

func (filter *InstanceFilter) VpcFilterExist() bool {
	return (len(filter.VpcName) > 0 || len(filter.VpcId) > 0)
}

func (filter *InstanceFilter) vpcDescribeParams() *ec2.DescribeVpcsInput {
	if len(filter.VpcId) > 0 {
		return &ec2.DescribeVpcsInput{
			VpcIds: []*string{
				aws.String(filter.VpcId),
			},
		}
	}

	if len(filter.VpcName) > 0 {
		return &ec2.DescribeVpcsInput{
			Filters: []*ec2.Filter{
				{
					Name: aws.String("tag:Name"),
					Values: []*string{
						aws.String(filter.VpcName),
					},
				},
			},
		}
	}

	return nil
}

func (filter *InstanceFilter) vpcIdForFilter() (*string, error) {
	params := filter.vpcDescribeParams()
	if params == nil {
		return nil, nil // without vpc filter
	}

	vpcs, err := findVpcs(params)
	if err != nil {
		return nil, err
	}

	return vpcs[0].VpcId, nil
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
