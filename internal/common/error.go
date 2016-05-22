package common

import (
	"fmt"
)

// VpcNotFoundError describe a error that  Vpc not found.
type VpcNotFoundError struct {
	Cause string
}

// Error return error message
func (err *VpcNotFoundError) Error() string {
	return fmt.Sprintf("No such vpc found. %s", err.Cause)
}
