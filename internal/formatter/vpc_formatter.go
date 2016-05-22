package formatter

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

// VpcFormatter describe a formatter to display VPC.
type VpcFormatter struct{}

// NewVpcFormatter create VpcFormatter instance.
func NewVpcFormatter() *VpcFormatter {
	return &VpcFormatter{}
}

func nameOfVpc(vpc *ec2.Vpc) string {
	if vpc == nil {
		return UNDEFINED
	}

	for _, t := range vpc.Tags {
		if *t.Key == "Name" {
			return *t.Value
		}
	}
	return UNDEFINED
}

// Format return formatted string which contains the VPC.
func (formatter *VpcFormatter) Format(vpc *ec2.Vpc) []string {
	return []string{
		nameOfVpc(vpc),
		*vpc.VpcId,
		*vpc.State,
	}
}
