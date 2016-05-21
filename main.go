package main

import (
	"os"

	"github.com/hirakiuc/ec2s/internal/command/elbs"
	"github.com/hirakiuc/ec2s/internal/command/list"
	"github.com/hirakiuc/ec2s/internal/command/scp"
	"github.com/hirakiuc/ec2s/internal/command/ssh"
	"github.com/hirakiuc/ec2s/internal/command/vpcs"
	"github.com/hirakiuc/ec2s/internal/common"

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
		"elbs": func() (cli.Command, error) {
			return elbs.GetCommand(), nil
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
