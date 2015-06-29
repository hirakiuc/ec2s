package list

import (
	"flag"
	"fmt"
	"os"

	"../../common"
	"../../config"
)

type Command struct {
	*common.InstanceFilter
}

func GetCommand() *Command {
	return &Command{
		&common.InstanceFilter{
			VpcId: "",
		},
	}
}

func (c *Command) Help() string {
	return "ec2s list"
}

func (c *Command) Run(args []string) int {
	c.parseOptions(args)
	return ShowEc2Instances(
		os.Stdout,
		common.InstancesFilter(c),
	)
}

func (c *Command) Synopsis() string {
	return "Show ec2 instances."
}

func (c *Command) parseOptions(args []string) {
	var configPath string

	f := flag.NewFlagSet("list", flag.ExitOnError)
	f.StringVar(&c.VpcId, "vpc-id", "", "vpc id")
	f.StringVar(&configPath, "c", "~/.ec2s.toml", "config path")
	f.Parse(args)

	_, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Can't load config file: %s, %v\n", configPath, err)
		os.Exit(1)
	}
}
