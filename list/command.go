package list

import (
	"bytes"
	"fmt"
	//	"os"
)

type Command struct{}

func (c *Command) Help() string {
	return "ec2s list"
}

func (c *Command) Run(args []string) int {
	// create []byte and writer to this buffer.
	//	output := os.Stdout
	buf := &bytes.Buffer{}
	showList(buf)
	fmt.Print(buf.String())
	return 0
}

func (c *Command) Synopsis() string {
	return "Show ec2 instances."
}
