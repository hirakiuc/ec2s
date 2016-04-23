package formatter

import (
	"github.com/hirakiuc/ec2s/cache"
)

const UNDEFINED = "---"

func vpcNameById(vpcId *string) string {
	if vpcId == nil {
		return UNDEFINED
	}

	vpcName := (cache.GetVpcCache()).VpcName(*vpcId)
	if vpcName != nil {
		return *vpcName
	} else {
		return UNDEFINED
	}
}
