package formatter

import (
	"github.com/aws/aws-sdk-go/service/elb"
)

type ElbFormatter struct{}

func NewElbFormatter() *ElbFormatter {
	return &ElbFormatter{}
}

func (formatter *ElbFormatter) Format(elb *elb.LoadBalancerDescription) []string {
	return []string{
		vpcNameById(elb.VPCId),
		*elb.LoadBalancerName,
		*elb.DNSName,
	}
}