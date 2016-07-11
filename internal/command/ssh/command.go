package ssh

import (
	"errors"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/hirakiuc/ec2s/internal/chooser"
	"github.com/hirakiuc/ec2s/internal/common"
	"github.com/hirakiuc/ec2s/internal/options"
)

// Command describe ssh command.
type Command struct {
	*common.InstanceFilter
	Command string
}

var logger *common.Logger
var command Command

func init() {
	logger = common.GetLogger()
	command = Command{
		InstanceFilter: &common.InstanceFilter{},
		Command:        "",
	}

	_, err := options.AddCommand(
		"ssh",
		"ssh to the selected ec2 instance",
		"ssh command invoke ssh to selected ec2 instance",
		&command)
	if err != nil {
		common.ShowError(err)
		os.Exit(1)
	}
}

// Execute invoke scp command.
func (c *Command) Execute(args []string) error {
	if err := c.validateOptions(args); err != nil {
		common.ShowError(err)
		return err
	}

	instances, err := chooser.ChooseEc2Instances(c)
	if err != nil {
		common.ShowError(err)
		return err
	}

	if len(instances) == 0 {
		return nil
	}

	if len(c.Command) > 0 {
		return c.execSSHCommand(instances)
	}

	return c.execSSHLogin(instances)
}

func (c *Command) validateOptions(args []string) error {
	opts := options.GetOptions()

	if err := opts.Validate(); err != nil {
		return err
	}

	if len(args) > 0 {
		c.Command = strings.Join(args, " ")
	}

	return nil
}

func (c *Command) execSSHLogin(instances []*ec2.Instance) error {
	if len(instances) > 1 {
		logger.Warn("ssh subcommand only use first selection.\n")
	}

	instance := instances[0]
	if err := common.IsNetworkAccessible(instance); err != nil {
		common.ShowError(err)
		return err
	}

	if err := c.execSSH(instances[0]); err != nil {
		common.ShowError(err)
		return err
	}

	return nil
}

func (c *Command) execSSHCommand(instances []*ec2.Instance) error {
	errored := false

	for _, instance := range instances {
		err := common.IsNetworkAccessible(instance)
		if err != nil {
			common.ShowError(err)
			errored = true
			continue
		}

		err = c.execSSH(instance)
		if err != nil {
			logger.Error("failed to ssh./n")
			common.ShowError(err)
			errored = true
		}
	}

	if errored {
		return errors.New("Some errors occurred.")
	}

	return nil
}
