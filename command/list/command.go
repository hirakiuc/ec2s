package list

import (
	"flag"
	"os"

	"../../common"
	"../../config"
)

type Command struct {
	*common.InstanceFilter
}

var logger *common.Logger

func init() {
	logger = common.GetLogger()
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
	return ShowEc2Instances(os.Stdout, c)
}

func (c *Command) Synopsis() string {
	return "Show ec2 instances."
}

func (c *Command) parseOptions(args []string) {
	var configPath string

	f := flag.NewFlagSet("list", flag.ExitOnError)
	f.StringVar(&c.VpcId, "vpc-id", "", "vpc id")
	f.StringVar(&c.VpcName, "vpc-name", "", "vpc name")
	f.StringVar(&configPath, "c", "~/.ec2s.toml", "config path")
	f.Parse(args)

	_, err := config.LoadConfig(configPath)
	if err != nil {
		logger.Error("Can't load config file: %s, %v\n", configPath, err)
		os.Exit(1)
	}
}
