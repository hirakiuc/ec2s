package main

import (
	"os"

	_ "github.com/hirakiuc/ec2s/internal/command/elbs"
	_ "github.com/hirakiuc/ec2s/internal/command/list"
	_ "github.com/hirakiuc/ec2s/internal/command/scp"
	_ "github.com/hirakiuc/ec2s/internal/command/ssh"
	_ "github.com/hirakiuc/ec2s/internal/command/version"
	_ "github.com/hirakiuc/ec2s/internal/command/vpcs"
	"github.com/hirakiuc/ec2s/internal/common"
	"github.com/hirakiuc/ec2s/internal/options"
)

var logger *common.Logger

func init() {
	logger = common.GetLogger()
}

func main() {
	if _, err := options.ParseOptions(); err != nil {
		os.Exit(1)
	}
}
