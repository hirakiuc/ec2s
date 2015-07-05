package scp

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"

	"../../config"
)

func (c *Command) execScp(instance *ec2.Instance) bool {
	conf := config.GetConfig()

	privateKeyPath := (conf.Ssh).IdentityFileForName(*instance.KeyName)
	if privateKeyPath == nil {
		fmt.Printf("Can't find private key Path: %s\n", *instance.KeyName)
		return false
	}

	cmd := exec.Command(
		"scp",
		"-P", strconv.Itoa(conf.Ssh.Port),
		"-i", *privateKeyPath,
		expandPath(c.FromPath, instance),
		expandPath(c.ToPath, instance),
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return false
	}

	cmd.Wait()
	return true
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

	prefix := fmt.Sprintf("%s@%s:", conf.Ssh.User, *instance.PublicIPAddress)
	return strings.Replace(path, "ec2:", prefix, 1)
}
