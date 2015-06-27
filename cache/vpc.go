package cache

import (
	"../common"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type VpcCache struct {
	Entries map[string]*ec2.VPC
}

var vpcCache *VpcCache

func GetVpcCache() *VpcCache {
	if vpcCache == nil {
		vpcCache = &VpcCache{
			Entries: map[string]*ec2.VPC{},
		}
	}

	return vpcCache
}

func (cache *VpcCache) ReadEntry(vpcId string) *ec2.VPC {
	return cache.Entries[vpcId]
}

func (cache *VpcCache) WriteEntry(vpcId string, vpc *ec2.VPC) {
	cache.Entries[vpcId] = vpc
}

func (cache *VpcCache) MakeCache() bool {
	service := common.Ec2Service()
	res, err := service.DescribeVPCs(nil)
	if err != nil {
		fmt.Println("failed to make vpcs cache.")
		common.ShowError(err)
		return false
	}

	for _, vpc := range res.VPCs {
		cache.WriteEntry(*vpc.VPCID, vpc)
	}
	return true
}
