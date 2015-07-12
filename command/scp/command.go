package scp

import (
	"flag"
	"fmt"
	"os"

	"../../chooser"
	"../../common"
	"../../config"
)

type Command struct {
	*common.InstanceFilter
	FromPath string
	ToPath   string
}

var logger *common.Logger

func init() {
	logger = common.GetLogger()
}

func GetCommand() *Command {
	return &Command{
		InstanceFilter: &common.InstanceFilter{
			VpcId:   "",
			VpcName: "",
		},
		FromPath: "",
		ToPath:   "",
	}
}

func (c *Command) Help() string {
	return "ec2s scp"
}

func (c *Command) Run(args []string) int {
	c.parseOptions(args)

	instances, ret := chooser.ChooseEc2Instances(c)
	if ret != 0 {
		return ret
	}

	if len(instances) == 0 {
		return 0
	}

	for _, instance := range instances {
		if common.IsNetworkAccessible(instance) == false {
			logger.Warn("%s is not reachable.\n", *instance.InstanceID)
		} else {
			c.execScp(instance)
		}
	}

	return 0
}

func (c *Command) Synopsis() string {
	return "scp from/to instance"
}

func (c *Command) parseOptions(args []string) {
	var configPath string

	f := flag.NewFlagSet("scp", flag.ExitOnError)
	f.StringVar(&c.VpcId, "vpc-id", "", "vpc id")
	f.StringVar(&c.VpcName, "vpc-name", "", "vpc name")
	f.StringVar(&configPath, "c", "~/.ec2s.toml", "config path")
	f.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		f.PrintDefaults()
	}
	f.Parse(args)

	conf, err := config.LoadConfig(configPath)
	if err != nil {
		logger.Error("Can't load config file: %s, %v\n", configPath, err)
		os.Exit(1)
	}

	logger := common.GetLogger()
	logger.SetColored(conf.Common.ColorizedOutput)

	if f.NArg() != 2 {
		f.Usage()
		os.Exit(1)
	}

	c.FromPath = f.Arg(0)
	c.ToPath = f.Arg(1)
}
