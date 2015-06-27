package cache

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Ec2InstanceCache struct {
	Entries map[string]*ec2.Instance
}

var ec2instanceCache *Ec2InstanceCache

func GetEc2InstanceCache() *Ec2InstanceCache {
	if ec2instanceCache == nil {
		ec2instanceCache = &Ec2InstanceCache{
			Entries: map[string]*ec2.Instance{},
		}
	}

	return ec2instanceCache
}

func (cache *Ec2InstanceCache) ReadEntry(instanceId string) *ec2.Instance {
	return cache.Entries[instanceId]
}

func (cache *Ec2InstanceCache) WriteEntry(instanceId string, instance *ec2.Instance) {
	cache.Entries[instanceId] = instance
}
