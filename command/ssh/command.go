package ssh

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/ec2"

	"../../common"
	"../../config"
)

type Command struct {
	*common.InstanceFilter
}

func GetCommand() *Command {
	return &Command{
		&common.InstanceFilter{
			VpcId: "",
		},
	}
}

func (c *Command) Help() string {
	return "ec2s ssh"
}

func (c *Command) Run(args []string) int {
	c.parseOptions(args)
	instances, ret := c.choseInstance()
	if ret != 0 {
		return ret
	}

	if len(instances) == 0 {
		return 0
	}
	if len(instances) > 1 {
		fmt.Println("WARN: ssh subcommand only use first selection.")
	}

	instance := instances[0]
	if isNetworkAccessible(instance) == false {
		return 1
	}

	if execSsh(instances[0]) == false {
		return 1
	} else {
		return 0
	}
}

func (c *Command) Synopsis() string {
	return "ssh to instance"
}

func (c *Command) parseOptions(args []string) {
	var configPath string

	f := flag.NewFlagSet("ssh", flag.ExitOnError)
	f.StringVar(&c.VpcId, "vpc-id", "", "vpc id")
	f.StringVar(&c.VpcName, "vpc-name", "", "vpc name")
	f.StringVar(&configPath, "c", "~/.ec2s.toml", "config path")
	f.Parse(args)

	_, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Can't load config file: %s, %v\n", configPath, err)
		os.Exit(1)
	}
}

func isNetworkAccessible(instance *ec2.Instance) bool {
	if *instance.State.Name != "running" {
		fmt.Printf("The instance is not running. (%s)\n", *instance.InstanceID)
		return false
	}

	if instance.PublicIPAddress == nil {
		fmt.Printf("The instance does not have Public IPAddress. (%s)\n", *instance.InstanceID)
		return false
	}

	return true
}
