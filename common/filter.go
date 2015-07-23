package common

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type InstanceFilter struct {
	VpcId       string
	VpcName     string
	Elbname     string
	InstanceIds []string
}

type FilterInterface interface {
	InstancesFilter() (*ec2.DescribeInstancesInput, error)
}

func (filter *InstanceFilter) vpcFilterExist() bool {
	return (len(filter.VpcName) > 0 || len(filter.VpcId) > 0)
}

func (filter *InstanceFilter) instanceIdsFilterExist() bool {
	return (len(filter.InstanceIds) > 0)
}

func (filter *InstanceFilter) vpcDescribeParams() *ec2.DescribeVPCsInput {
	if len(filter.VpcId) > 0 {
		return &ec2.DescribeVPCsInput{
			VPCIDs: []*string{
				aws.String(filter.VpcId),
			},
		}
	}

	if len(filter.VpcName) > 0 {
		return &ec2.DescribeVPCsInput{
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

	return vpcs[0].VPCID, nil
}

func (filter *InstanceFilter) vpcFilter() (*ec2.Filter, error) {
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
	filters := []*ec2.Filter{}

	vpcFilter, err := filter.vpcFilter()
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
