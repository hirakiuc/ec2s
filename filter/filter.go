package filter

import (
	"../common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Filter struct {
	VpcId       string
	VpcName     string
	ElbName     string
	InstanceIds []string
}

type FilterInterface interface {
	InstancesFilter() (*ec2.DescribeInstancesInput, error)
}

func (filter *Filter) vpcFilterExist() bool {
	return (len(filter.VpcName) > 0 || len(filter.VpcId) > 0)
}

func (filter *Filter) instanceIdsFilterExist() bool {
	return (len(filter.InstanceIds) > 0)
}

func (filter *Filter) vpcDescribeParams() *ec2.DescribeVPCsInput {
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

func (filter *Filter) vpcIdForFilter() (*string, error) {
	params := filter.vpcDescribeParams()
	if params == nil {
		return nil, nil // without vpc filter
	}

	vpcs, err := common.FindVpcs(params)
	if err != nil {
		return nil, err
	}

	return vpcs[0].VPCID, nil
}

func (filter *Filter) vpcFilter() (*ec2.Filter, error) {
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

func (filter *Filter) InstancesFilter() (*ec2.DescribeInstancesInput, error) {
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
