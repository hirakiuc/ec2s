package list

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
	return "ec2s list"
}

func (c *Command) Run(args []string) int {
	c.parseOptions(args)
	return ShowEc2Instances(os.Stdout)
}

func (c *Command) Synopsis() string {
	return "Show ec2 instances."
}

func (c *Command) parseOptions(args []string) {
	var configPath string

	f := flag.NewFlagSet("list", flag.ExitOnError)
	f.StringVar(&c.VpcName, "vpc-name", "", "vpc name")
	f.StringVar(&configPath, "c", "~/.ec2s.toml", "config path")
	f.Parse(args)

	fmt.Println(c.VpcName)

	_, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Can't load config file: %s, %v\n", configPath, err)
		os.Exit(1)
	}
}
