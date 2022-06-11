package cache

import (
	"github.com/hirakiuc/ec2s/internal/common"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// VpcCache has Cached Vpcs
type VpcCache struct {
	Entries map[string]*ec2.Vpc
}

var (
	vpcCache *VpcCache
	logger   *common.Logger
)

func init() {
	logger = common.GetLogger()
}

// GetVpcCache return the VpcCache instance.
func GetVpcCache() *VpcCache {
	if vpcCache == nil {
		vpcCache = &VpcCache{
			Entries: map[string]*ec2.Vpc{},
		}
	}

	return vpcCache
}

// ReadEntry return a cached Vpc which identifiied by vpcID
func (cache *VpcCache) ReadEntry(vpcID string) *ec2.Vpc {
	return cache.Entries[vpcID]
}

// WriteEntry cache the specified Vpc in VpcCache.
func (cache *VpcCache) WriteEntry(vpcID string, vpc *ec2.Vpc) {
	cache.Entries[vpcID] = vpc
}

// MakeCache create Vpc Cache
func (cache *VpcCache) MakeCache() error {
	service, err := common.Ec2Service()
	if err != nil {
		return err
	}

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

// VpcName return name of the vpc which specified by vpcID
func (cache *VpcCache) VpcName(vpcID string) *string {
	vpc := cache.ReadEntry(vpcID)
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
