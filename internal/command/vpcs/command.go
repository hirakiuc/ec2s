package vpcs

import (
	"os"

	"github.com/hirakiuc/ec2s/internal/common"
	"github.com/hirakiuc/ec2s/internal/options"
)

// Command describe vpcs command.
type Command struct{}

var logger *common.Logger
var command Command

func init() {
	logger = common.GetLogger()
	command = Command{}

	_, err := options.AddCommand(
		"vpcs",
		"List vpc definitions",
		"vpcs command show vpc definitions.",
		&command)
	if err != nil {
		common.ShowError(err)
		os.Exit(1)
	}
}

// Execute invoke vpcs command.
func (c *Command) Execute(args []string) error {
	if err := c.validateOptions(args); err != nil {
		common.ShowError(err)
		return err
	}

	if err := c.showVpcs(os.Stdout); err != nil {
		common.ShowError(err)
		return err
	}

	return nil
}

func (c *Command) validateOptions(args []string) error {
	opts := options.GetOptions()

	if err := opts.Validate(); err != nil {
		return err
	}

	return nil
}
