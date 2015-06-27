package formatter

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

const UNDEFINED = "---"

type Ec2Formatter struct{}

func NewEc2Formatter() *Ec2Formatter {
	return &Ec2Formatter{}
}

func nameOfInstance(instance *ec2.Instance) string {
	for _, t := range instance.Tags {
		if *t.Key == "Name" {
			return *t.Value
		}
	}
	return UNDEFINED
}

func vpcName(vpc *ec2.VPC) string {
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

func ipAddress(instance *ec2.Instance) string {
	if instance.PublicIPAddress != nil {
		return *instance.PublicIPAddress
	} else {
		return UNDEFINED
	}
}

func (formatter *Ec2Formatter) Format(vpc *ec2.VPC, instance *ec2.Instance) []string {
	return []string{
		vpcName(vpc),
		nameOfInstance(instance),
		*instance.InstanceID,
		*instance.InstanceType,
		*instance.Placement.AvailabilityZone,
		ipAddress(instance),
		*instance.State.Name,
	}
}
