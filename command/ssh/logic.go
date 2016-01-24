package ssh

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/aws/aws-sdk-go/service/ec2"

	"../../common"
	"../../config"
)

func (c *Command) logCommand(instance *ec2.Instance, privateKeyPath *string) {
	logger := common.GetLogger()
	conf := config.GetConfig()

	logger.Info("ssh -l %s -p %d -i %s %s %s\n",
		conf.Ssh.User,
		conf.Ssh.Port,
		*privateKeyPath,
		*instance.PublicIpAddress,
		c.Command,
	)
}

func (c *Command) execSsh(instance *ec2.Instance) error {
	conf := config.GetConfig()

	privateKeyPath, err := (conf.Ssh).IdentityFileForName(*instance.KeyName)
	if err != nil {
		return err
	}

	c.logCommand(instance, privateKeyPath)

	cmd := exec.Command(
		"ssh",
		"-l", conf.Ssh.User,
		"-p", strconv.Itoa(conf.Ssh.Port),
		"-i", *privateKeyPath,
		*instance.PublicIpAddress,
		c.Command,
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		logger.Error("failed to execute command.\n", err)
		return err
	}

	cmd.Wait()
	return nil
}
