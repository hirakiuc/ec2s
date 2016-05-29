package common

import (
	"fmt"
)

// ArgumentError describe a error that receiver passed invalid arguments.
type ArgumentError struct {
	Cause string
}

// NewArgumentError create ArgumentError instance.
func NewArgumentError(msg string) ArgumentError {
	return ArgumentError{
		Cause: msg,
	}
}

// Error return error message
func (err ArgumentError) Error() string {
	return fmt.Sprintf("Argument Error: %s", err.Cause)
}

// VpcNotFoundError describe a error that  Vpc not found.
type VpcNotFoundError struct {
	Cause string
}

// Error return error message
func (err VpcNotFoundError) Error() string {
	return fmt.Sprintf("No such vpc found. %s", err.Cause)
}
