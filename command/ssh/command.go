package ssh

import (
	"flag"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"

	"../../chooser"
	"../../common"
	"../../config"
)

type Command struct {
	*common.InstanceFilter
	Command string
}

var logger *common.Logger

func init() {
	logger = common.GetLogger()
}

func GetCommand() *Command {
	return &Command{
		InstanceFilter: &common.InstanceFilter{
			VpcId: "",
		},
		Command: "",
	}
}

func (c *Command) Help() string {
	return "ec2s ssh"
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

	if len(c.Command) > 0 {
		return c.execSshCommand(instances)
	} else {
		return c.execSshLogin(instances)
	}

	if len(instances) > 1 {
		logger.Warn("ssh subcommand only use first selection.\n")
	}

	instance := instances[0]
	if common.IsNetworkAccessible(instance) == false {
		return 1
	}

	if err := c.execSsh(instances[0]); err != nil {
		common.ShowError(err)
		return 1
	} else {
		return 0
	}
}

func (c *Command) Synopsis() string {
	return "ssh to instance"
}

func (c *Command) parseOptions(args []string) error {
	var configPath string

	f := flag.NewFlagSet("ssh", flag.ExitOnError)
	f.StringVar(&c.VpcId, "vpc-id", "", "vpc id")
	f.StringVar(&c.VpcName, "vpc-name", "", "vpc name")
	f.StringVar(&configPath, "c", "~/.ec2s.toml", "config path")
	f.Parse(args)

	conf, err := config.LoadConfig(configPath)
	if err != nil {
		logger.Error("Can't load config file.\n")
		return err
	}

	logger := common.GetLogger()
	logger.SetColored(conf.Common.ColorizedOutput)

	if f.NArg() > 0 {
		c.Command = strings.Join(f.Args(), " ")
	}

	return nil
}

func (c *Command) execSshLogin(instances []*ec2.Instance) int {
	if len(instances) > 1 {
		logger.Warn("ssh subcommand only use first selection.\n")
	}

	instance := instances[0]
	if common.IsNetworkAccessible(instance) == false {
		return 1
	}

	if err := c.execSsh(instances[0]); err != nil {
		common.ShowError(err)
		return 1
	} else {
		return 0
	}
}

func (c *Command) execSshCommand(instances []*ec2.Instance) int {
	ret := 0

	for _, instance := range instances {
		if common.IsNetworkAccessible(instance) == true {
			if err := c.execSsh(instance); err != nil {
				common.ShowError(err)
				ret = ret + 1
			}
		}
	}

	return ret
}
