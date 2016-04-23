package chooser

import (
	"bytes"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"
	"gopkg.in/pipe.v2"

	"github.com/hirakiuc/ec2s/cache"
	"github.com/hirakiuc/ec2s/command/list"
	"github.com/hirakiuc/ec2s/common"
	"github.com/hirakiuc/ec2s/config"
)

var logger *common.Logger

func init() {
	logger = common.GetLogger()
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
		logger.Error("Can't find ec2 instance: '%s'\n", instanceId)
	}

	return instance
}

func ChooseEc2Instances(options common.FilterInterface) ([]*ec2.Instance, error) {
	buffer := bytes.NewBuffer(nil)
	err := list.ShowEc2Instances(buffer, options)
	if err != nil {
		return []*ec2.Instance{}, err
	}

	conf := config.GetConfig()

	p := pipe.Line(
		pipe.Print(buffer.String()),
		pipe.Exec(conf.Peco.Path),
	)

	output, err := pipe.CombinedOutput(p)
	if err != nil {
		logger.Error("Command failed.\n", err)
		return []*ec2.Instance{}, err
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

	return instances, nil
}
