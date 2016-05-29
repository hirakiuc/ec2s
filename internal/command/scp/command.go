package scp

import (
	"flag"
	"fmt"
	"os"

	"github.com/hirakiuc/ec2s/internal/chooser"
	"github.com/hirakiuc/ec2s/internal/common"
	"github.com/hirakiuc/ec2s/internal/config"
)

// Command describe sccp command.
type Command struct {
	*common.InstanceFilter
	FromPath string
	ToPath   string
}

var logger *common.Logger

func init() {
	logger = common.GetLogger()
}

// GetCommand create scp command instance.
func GetCommand() *Command {
	return &Command{
		InstanceFilter: &common.InstanceFilter{
			VpcID:   "",
			VpcName: "",
		},
		FromPath: "",
		ToPath:   "",
	}
}

// Help return help message.
func (c *Command) Help() string {
	return "ec2s scp"
}

// Run invoke scp command.
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
				cnt++
			}
		}
	}

	return cnt
}

// Synopsis return command description.
func (c *Command) Synopsis() string {
	return "scp from/to instance"
}

func (c *Command) parseOptions(args []string) error {
	var configPath string

	f := flag.NewFlagSet("scp", flag.ExitOnError)
	f.StringVar(&c.VpcID, "vpc-id", "", "vpc id")
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
		err := common.NewArgumentError("Require two arguments.")
		common.ShowError(err)
		f.Usage()
		return err
	}

	c.FromPath = f.Arg(0)
	c.ToPath = f.Arg(1)

	return nil
}
