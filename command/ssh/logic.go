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
		*instance.PublicIPAddress,
		c.Command,
	)
}

func (c *Command) execSsh(instance *ec2.Instance) bool {
	conf := config.GetConfig()

	privateKeyPath := (conf.Ssh).IdentityFileForName(*instance.KeyName)
	if privateKeyPath == nil {
		logger.Error("Can't find private key Path: %s\n", *instance.KeyName)
		return false
	}

	c.logCommand(instance, privateKeyPath)

	cmd := exec.Command(
		"ssh",
		"-l", conf.Ssh.User,
		"-p", strconv.Itoa(conf.Ssh.Port),
		"-i", *privateKeyPath,
		*instance.PublicIPAddress,
		c.Command,
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		logger.Error("failed to execute command: %v\n", err)
		return false
	}

	cmd.Wait()
	return true
}
