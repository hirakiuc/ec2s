package list

import (
	"os"

	"github.com/hirakiuc/ec2s/internal/common"
	"github.com/hirakiuc/ec2s/internal/options"
)

// Command describe list command.
type Command struct {
	*common.InstanceFilter
}

var logger *common.Logger
var command Command

func init() {
	logger = common.GetLogger()
	command = Command{&common.InstanceFilter{}}

	_, err := options.AddCommand(
		"list",
		"List ec2 instances.",
		"list command show ec2 instances.",
		&command)
	if err != nil {
		logger.Error("Internal Error: %v", err)
		os.Exit(1)
	}
}

// Execute invoke list command.
func (c *Command) Execute(args []string) error {
	if err := c.validateOptions(args); err != nil {
		common.ShowError(err)
		return err
	}

	if err := ShowEc2Instances(os.Stdout, c); err != nil {
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
