package formatter

import (
	"github.com/aws/aws-sdk-go/service/elb"
)

// ElbFormatter describe a formatter to display ELB.
type ElbFormatter struct{}

// NewElbFormatter create ElbFormatter instance.
func NewElbFormatter() *ElbFormatter {
	return &ElbFormatter{}
}

// Format return formatted string which contains the ELB.
func (formatter *ElbFormatter) Format(elb *elb.LoadBalancerDescription) []string {
	return []string{
		vpcNameByID(elb.VPCId),
		*elb.LoadBalancerName,
		*elb.DNSName,
	}
}
