package main

import (
	"os"

	"./command/list"
	"./command/scp"
	"./command/ssh"
	"./command/vpcs"
	"./common"

	"github.com/mitchellh/cli"
)

const VERSION string = "0.1.0"

var logger *common.Logger

func init() {
	logger = common.GetLogger()
}

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
		logger.Error(err.Error())
	}
	os.Exit(exitStatus)
}
