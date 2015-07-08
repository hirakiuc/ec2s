package main

import (
	"fmt"
	"os"

	"./command/list"
	"./command/scp"
	"./command/ssh"
	"./command/vpcs"

	"github.com/mitchellh/cli"
)

const VERSION string = "0.1.0"

func main() {
	c := cli.NewCLI("ec2s", VERSION)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return list.GetCommand(), nil
		},
		"vpcs": func() (cli.Command, error) {
			return vpcs.GetCommand(), nil
		},
		"ssh": func() (cli.Command, error) {
			return ssh.GetCommand(), nil
		},
		"scp": func() (cli.Command, error) {
			return scp.GetCommand(), nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(exitStatus)
}
