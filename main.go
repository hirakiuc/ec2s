package main

import (
	"fmt"
	"os"

	"./list"
	"github.com/mitchellh/cli"
)

const VERSION string = "0.1.0"

func main() {
	c := cli.NewCLI("ec2s", VERSION)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &list.Command{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(exitStatus)
}
