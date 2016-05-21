package common

import (
	"fmt"
)

type ErrorInterface interface {
}

type VpcNotFoundError struct {
	Cause string
}

// function to implement 'type error interface'
func (err *VpcNotFoundError) Error() string {
	return fmt.Sprintf("No such vpc found. %s", err.Cause)
}
