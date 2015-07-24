package scp

import (
	"flag"
	"fmt"
	"os"

	"../../chooser"
	"../../common"
	"../../config"
	"../../filter"
)

type Command struct {
	*filter.Filter
	FromPath string
	ToPath   string
}

var logger *common.Logger

func init() {
	logger = common.GetLogger()
}

func GetCommand() *Command {
	return &Command{
		Filter: &filter.Filter{
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
	if err := c.parseOptions(args); err != nil {
		common.ShowError(err)
		return 1
	}

	instances, err := chooser.ChooseEc2Instances(c)
	if err != nil {
		common.ShowError(err)
		return 1
	}

	if len(instances) == 0 {
		return 0
	}

	cnt := 0
	for _, instance := range instances {
		if common.IsNetworkAccessible(instance) == true {
			if err := c.execScp(instance); err != nil {
				logger.Error("failed to scp.\n")
				common.ShowError(err)
				cnt += 1
			}
		}
	}

	return cnt
}

func (c *Command) Synopsis() string {
	return "scp from/to instance"
}

func (c *Command) parseOptions(args []string) error {
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
		logger.Error("Can't load config file.\n")
		return err
	}

	logger := common.GetLogger()
	logger.SetColored(conf.Common.ColorizedOutput)

	if f.NArg() != 2 {
		f.Usage()
		os.Exit(1) // TODO: fix to return error.
	}

	c.FromPath = f.Arg(0)
	c.ToPath = f.Arg(1)

	return nil
}
