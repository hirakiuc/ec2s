package formatter

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Ec2Formatter describe a formatter to display EC2 instance.
type Ec2Formatter struct{}

// NewEc2Formatter create Ec2Formatter instance.
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

func ipAddress(instance *ec2.Instance) string {
	if instance.PublicIpAddress != nil {
		return *instance.PublicIpAddress
	}

	return UNDEFINED
}

// Format return formatted string which contains the EC2 instance.
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
