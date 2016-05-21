package formatter

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

type VpcFormatter struct{}

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

func (formatter *VpcFormatter) Format(vpc *ec2.Vpc) []string {
	return []string{
		nameOfVpc(vpc),
		*vpc.VpcId,
		*vpc.State,
	}
}