package ssh

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/hirakiuc/ec2s/internal/common"
	"github.com/hirakiuc/ec2s/internal/config"
)

func (c *Command) logCommand(instance *ec2.Instance, privateKeyPath *string) {
	logger := common.GetLogger()
	conf := config.GetConfig()

	logger.Info("ssh -l %s -p %d -i %s %s %s\n",
		conf.SSH.User,
		conf.SSH.Port,
		*privateKeyPath,
		*instance.PublicIpAddress,
		c.Command,
	)
}

func (c *Command) execSSH(instance *ec2.Instance) error {
	conf := config.GetConfig()

	privateKeyPath, err := (conf.SSH).IdentityFileForName(*instance.KeyName)
	if err != nil {
		return err
	}

	c.logCommand(instance, privateKeyPath)

	cmd := exec.Command(
		"ssh",
		"-l", conf.SSH.User,
		"-p", strconv.Itoa(conf.SSH.Port),
		"-i", *privateKeyPath,
		*instance.PublicIpAddress,
		c.Command,
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		logger.Error("failed to execute command: %v\n", err)
		return err
	}

	return cmd.Wait()
}
