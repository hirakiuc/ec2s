package formatter

import (
	"github.com/hirakiuc/ec2s/internal/cache"
)

// UNDEFINED define the string to show 'the value is not defined.'.
const UNDEFINED = "---"

func vpcNameByID(vpcID *string) string {
	if vpcID == nil {
		return UNDEFINED
	}

	vpcName := (cache.GetVpcCache()).VpcName(*vpcID)
	if vpcName != nil {
		return *vpcName
	}

	return UNDEFINED
}
