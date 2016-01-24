package formatter

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

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

func vpcName(vpc *ec2.Vpc) string {
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
	if instance.PublicIpAddress != nil {
		return *instance.PublicIpAddress
	} else {
		return UNDEFINED
	}
}

func (formatter *Ec2Formatter) Format(vpc *ec2.Vpc, instance *ec2.Instance) []string {
	return []string{
		nameOfVpc(vpc),
		nameOfInstance(instance),
		*instance.InstanceId,
		*instance.InstanceType,
		*instance.Placement.AvailabilityZone,
		ipAddress(instance),
		*instance.State.Name,
	}
}
