package cache

import (
	"github.com/hirakiuc/ec2s/common"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type VpcCache struct {
	Entries map[string]*ec2.Vpc
}

var vpcCache *VpcCache
var logger *common.Logger

func init() {
	logger = common.GetLogger()
}

func GetVpcCache() *VpcCache {
	if vpcCache == nil {
		vpcCache = &VpcCache{
			Entries: map[string]*ec2.Vpc{},
		}
	}

	return vpcCache
}

func (cache *VpcCache) ReadEntry(vpcId string) *ec2.Vpc {
	return cache.Entries[vpcId]
}

func (cache *VpcCache) WriteEntry(vpcId string, vpc *ec2.Vpc) {
	cache.Entries[vpcId] = vpc
}

func (cache *VpcCache) MakeCache() error {
	service := common.Ec2Service()
	res, err := service.DescribeVpcs(nil)
	if err != nil {
		logger.Error("failed to make vpcs cache.\n")
		return err
	}

	for _, vpc := range res.Vpcs {
		cache.WriteEntry(*vpc.VpcId, vpc)
	}
	return nil
}

func (cache *VpcCache) VpcName(vpcId string) *string {
	vpc := cache.ReadEntry(vpcId)
	if vpc == nil {
		return nil
	}

	for _, t := range vpc.Tags {
		if *t.Key == "Name" {
			return t.Value
		}
	}

	return nil
}
