package chooser

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"
	"gopkg.in/pipe.v2"

	"../cache"
	"../command/list"
	"../common"
	"../config"
)

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

// TODO return error, not int.
func ChooseEc2Instances(options common.FilterInterface) ([]*ec2.Instance, int) {
	buffer := bytes.NewBuffer(nil)
	ret := list.ShowEc2Instances(buffer, options)
	if ret != 0 {
		return []*ec2.Instance{}, ret
	}

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
