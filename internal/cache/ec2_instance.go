package cache

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Ec2InstanceCache has Cached Ec2Instances
type Ec2InstanceCache struct {
	Entries map[string]*ec2.Instance
}

var ec2instanceCache *Ec2InstanceCache

// GetEc2InstanceCache return the Ec2InstanceCache instance
func GetEc2InstanceCache() *Ec2InstanceCache {
	if ec2instanceCache == nil {
		ec2instanceCache = &Ec2InstanceCache{
			Entries: map[string]*ec2.Instance{},
		}
	}

	return ec2instanceCache
}

// ReadEntry return a cached Ec2Instance which identified by instanceID.
func (cache *Ec2InstanceCache) ReadEntry(instanceID string) *ec2.Instance {
	return cache.Entries[instanceID]
}

// WriteEntry cache the specified ec2Instance in Ec2InstanceCache.
func (cache *Ec2InstanceCache) WriteEntry(instanceID string, instance *ec2.Instance) {
	cache.Entries[instanceID] = instance
}
