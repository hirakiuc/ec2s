package common

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// InstanceFilter is a filter condition about EC2s.
type InstanceFilter struct {
	VpcId   string
	VpcName string
}

// FilterInterface define interface of filter objects.
type FilterInterface interface {
	VpcFilter() (*ec2.Filter, error)
	InstancesFilter() (*ec2.DescribeInstancesInput, error)
	VpcFilterExist() bool
}

// VpcFilterExist check whether the filter contains vpc condition or not.
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

func (filter *InstanceFilter) vpcIDForFilter() (*string, error) {
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

// VpcFilter create ec2.Filter instance.
func (filter *InstanceFilter) VpcFilter() (*ec2.Filter, error) {
	vpcID, err := filter.vpcIDForFilter()
	if err != nil {
		return nil, err
	}

	if vpcID == nil {
		return nil, nil
	}

	return &ec2.Filter{
		Name: aws.String("vpc-id"),
		Values: []*string{
			aws.String(*vpcID),
		},
	}, nil
}

// InstancesFilter create ec2.DescribeInstancesInput instance.
func (filter *InstanceFilter) InstancesFilter() (*ec2.DescribeInstancesInput, error) {
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
	}

	return &ec2.DescribeInstancesInput{}, nil
}
