package formatter

import (
	"../cache"
)

const UNDEFINED = "---"

func vpcName(vpcId *string) string {
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
