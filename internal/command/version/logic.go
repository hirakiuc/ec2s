package version

import (
	"fmt"
	"io"
)

// ShowVersion print VERSION.
func (c *Command) ShowVersion(writer io.Writer) error {
	fmt.Fprintf(writer, "%s\n", VERSION)
	return nil
}
