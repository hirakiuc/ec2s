package vpcs

import (
	"flag"
	"fmt"
	"os"

	"../../config"
)

type Command struct {
	VpcName string
	VpcId   string
}

func (c *Command) Help() string {
	return "ec2s vpcs"
}

func (c *Command) Run(args []string) int {
	c.parseOptions(args)
	return c.showVpcs(os.Stdout)
}

func (c *Command) Synopsis() string {
	return "Show vpcs."
}

func (c *Command) parseOptions(args []string) {
	var configPath string
	f := flag.NewFlagSet("vpcs", flag.ExitOnError)
	f.StringVar(&configPath, "c", "~/.ec2s.toml", "config path")
	f.Parse(args)

	_, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Can't load config file: %s, %v\n", configPath, err)
		os.Exit(1)
	}
}
