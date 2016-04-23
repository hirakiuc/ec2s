package list

import (
	"flag"
	"os"

	"github.com/hirakiuc/ec2s/common"
	"github.com/hirakiuc/ec2s/config"
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
	if err := c.parseOptions(args); err != nil {
		common.ShowError(err)
		return 1
	}

	if err := ShowEc2Instances(os.Stdout, c); err != nil {
		common.ShowError(err)
		return 1
	}

	return 0
}

func (c *Command) Synopsis() string {
	return "Show ec2 instances."
}

func (c *Command) parseOptions(args []string) error {
	var configPath string

	f := flag.NewFlagSet("list", flag.ExitOnError)
	f.StringVar(&c.VpcId, "vpc-id", "", "vpc id")
	f.StringVar(&c.VpcName, "vpc-name", "", "vpc name")
	f.StringVar(&configPath, "c", "~/.ec2s.toml", "config path")
	f.Parse(args)

	_, err := config.LoadConfig(configPath)
	if err != nil {
		logger.Error("Can't load config file.\n")
		return err
	}

	return nil
}
