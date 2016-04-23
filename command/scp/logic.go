package scp

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/hirakiuc/ec2s/common"
	"github.com/hirakiuc/ec2s/config"
)

func (c *Command) logCommand(instance *ec2.Instance, privateKeyPath *string, fromPath string, toPath string) {
	logger := common.GetLogger()
	conf := config.GetConfig()

	logger.Info("scp -P %d -i %s %s %s\n",
		conf.Ssh.Port,
		*privateKeyPath,
		fromPath,
		toPath,
	)
}

func (c *Command) execScp(instance *ec2.Instance) error {
	conf := config.GetConfig()

	privateKeyPath, err := (conf.Ssh).IdentityFileForName(*instance.KeyName)
	if err != nil {
		return err
	}

	fromPath := expandPath(c.FromPath, instance)
	toPath := expandPath(c.ToPath, instance)

	c.logCommand(instance, privateKeyPath, fromPath, toPath)

	cmd := exec.Command(
		"scp",
		"-P", strconv.Itoa(conf.Ssh.Port),
		"-i", *privateKeyPath,
		fromPath,
		toPath,
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		logger.Error("failed to execute command.\n")
		return err
	}

	cmd.Wait()
	return nil
}

func expandPath(path string, instance *ec2.Instance) string {
	return expandRemotePath(
		expandTilda(path),
		instance,
	)
}

func expandTilda(path string) string {
	if len(path) < 2 {
		return path
	}

	if path[:2] != "~/" {
		return path
	}

	user, _ := user.Current()
	return strings.Replace(path, "~/", user.HomeDir+string(os.PathSeparator), 1)
}

func expandRemotePath(path string, instance *ec2.Instance) string {
	if len(path) < 4 {
		return path
	}

	if path[:4] != "ec2:" {
		return path
	}

	// expand 'ec2:' => 'user@ipaddr:'
	conf := config.GetConfig()

	prefix := fmt.Sprintf("%s@%s:", conf.Ssh.User, *instance.PublicIpAddress)
	return strings.Replace(path, "ec2:", prefix, 1)
}
