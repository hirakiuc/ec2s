package vpcs

import (
	"flag"
	"os"

	"github.com/hirakiuc/ec2s/internal/common"
	"github.com/hirakiuc/ec2s/internal/config"
)

type Command struct{}

var logger *common.Logger

func init() {
	logger = common.GetLogger()
}

func GetCommand() *Command {
	return &Command{}
}

func (c *Command) Help() string {
	return "ec2s vpcs"
}

func (c *Command) Run(args []string) int {
	if err := c.parseOptions(args); err != nil {
		common.ShowError(err)
		return 1
	}

	if err := c.showVpcs(os.Stdout); err != nil {
		common.ShowError(err)
		return 1
	}

	return 0
}

func (c *Command) Synopsis() string {
	return "Show vpcs."
}

func (c *Command) parseOptions(args []string) error {
	var configPath string
	f := flag.NewFlagSet("vpcs", flag.ExitOnError)
	f.StringVar(&configPath, "c", "~/.ec2s.toml", "config path")
	f.Parse(args)

	_, err := config.LoadConfig(configPath)
	if err != nil {
		logger.Error("Can't load config file.\n")
		return err
	}

	return nil
}
