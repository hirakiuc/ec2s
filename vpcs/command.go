package vpcs

import (
	"../config"
	"flag"
	"fmt"
	"os"
)

type Command struct {
	VpcName string
	VpcId   string
}

func (c *Command) Help() string {
	return "ec2s vpcs"
}

func (c *Command) Run(args []string) int {
	return 0
}
