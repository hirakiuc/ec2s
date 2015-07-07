package ssh

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/aws/aws-sdk-go/service/ec2"

	"../../config"
)

func execSsh(instance *ec2.Instance) bool {
	conf := config.GetConfig()

	privateKeyPath := (conf.Ssh).IdentityFileForName(*instance.KeyName)
	if privateKeyPath == nil {
		fmt.Printf("Can't find private key Path: %s\n", *instance.KeyName)
		return false
	}

	cmd := exec.Command(
		"ssh",
		"-l", conf.Ssh.User,
		"-p", strconv.Itoa(conf.Ssh.Port),
		"-i", *privateKeyPath,
		*instance.PublicIPAddress,
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
