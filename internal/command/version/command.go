package version

import (
	"os"

	"github.com/hirakiuc/ec2s/internal/common"
	"github.com/hirakiuc/ec2s/internal/options"
)

// VERSION number
const VERSION string = "0.1.0"

// Command describe version command.
type Command struct{}

var logger *common.Logger
var command Command

func init() {
	logger = common.GetLogger()
	command = Command{}

	_, err := options.AddCommand(
		"version",
		"show version",
		"version command show this tool version",
		&command)
	if err != nil {
		common.ShowError(err)
		os.Exit(1)
	}
}

// Execute invoke version command.
func (c *Command) Execute(args []string) error {
	if err := c.validateOptions(args); err != nil {
		common.ShowError(err)
		return err
	}

	if err := c.ShowVersion(os.Stdout); err != nil {
		common.ShowError(err)
		return err
	}

	return nil
}

func (c *Command) validateOptions(args []string) error {
	return nil
}
