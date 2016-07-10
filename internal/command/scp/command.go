package scp

import (
	"errors"
	"os"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hirakiuc/ec2s/internal/chooser"
	"github.com/hirakiuc/ec2s/internal/common"
	"github.com/hirakiuc/ec2s/internal/options"
)

// Command describe sccp command.
type Command struct {
	*common.InstanceFilter
	FromPath string
	ToPath   string
}

var logger *common.Logger
var command Command

func init() {
	logger = common.GetLogger()
	command = Command{
		InstanceFilter: &common.InstanceFilter{},
		FromPath:       "",
		ToPath:         "",
	}

	_, err := options.AddCommand(
		"scp",
		"scp from/to selected ec2 instances",
		"scp command invoke scp to selected ec2 instances",
		&command)
	if err != nil {
		logger.Error("Internal Error: %v", err)
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

	return c.execScpWithInstances(instances)
}

func (c *Command) execScpWithInstances(instances []*ec2.Instance) error {
	errored := false

	for _, instance := range instances {
		err := common.IsNetworkAccessible(instance)
		if err != nil {
			common.ShowError(err)
			errored = true
			continue
		}

		err = c.execScp(instance)
		if err != nil {
			logger.Error("failed to scp.\n")
			common.ShowError(err)
			errored = true
		}
	}

	if errored {
		return errors.New("Some errors occurred.")
	}

	return nil
}

func (c *Command) validateOptions(args []string) error {
	opts := options.GetOptions()

	if err := opts.Validate(); err != nil {
		return err
	}

	if len(args) != 2 {
		err := common.NewArgumentError("Require two arguments.")
		common.ShowError(err)
		return err
	}

	c.FromPath = args[0]
	c.ToPath = args[1]

	return nil
}
