package ssh

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"
	"gopkg.in/pipe.v2"

	"../../cache"
	"../../config"
	"../list"
)

func listEc2Instances(writer io.Writer, filter *ec2.DescribeInstancesInput) int {
	return list.ShowEc2Instances(writer, filter)
}

func ec2instance(line string) *ec2.Instance {
	if len(line) == 0 {
		return nil
	}

	vars := strings.Split(line, "\t")

	instanceId := strings.Trim(vars[2], " ")

	cache := cache.GetEc2InstanceCache()
	instance := cache.ReadEntry(instanceId)
	if instance == nil {
		fmt.Printf("Can't find ec2 instance: '%s'\n", instanceId)
	}

	return instance
}

func (c *Command) choseInstance(filter *ec2.DescribeInstancesInput) ([]*ec2.Instance, int) {
	buffer := bytes.NewBuffer(nil)
	listEc2Instances(buffer, filter)

	conf := config.GetConfig()

	p := pipe.Line(
		pipe.Print(buffer.String()),
		pipe.Exec(conf.Peco.Path),
	)
	output, err := pipe.CombinedOutput(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return []*ec2.Instance{}, 1
	}

	// parse line
	lines := strings.Split(string(output), "\n")

	instances := []*ec2.Instance{}
	for _, line := range lines {
		instance := ec2instance(line)
		if instance != nil {
			instances = append(instances, instance)
		}
	}

	return instances, 0
}

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
