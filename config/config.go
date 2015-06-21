package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"os/user"
	"strings"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

type Config struct {
	Aws Aws
}

type Aws struct {
	AccessKeyId     string `toml:"ACCESS_KEY_ID"`
	SecretAccessKey string `toml:"SECRET_ACCESS_KEY"`
	Region          string `toml:"REGION"`
}

var config Config

func GetConfig() *Config {
	return &config
}

func LoadConfig(path string) (*Config, error) {
	_, err := toml.DecodeFile(expandPath(path), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func expandPath(path string) string {
	if path[:2] == "~/" {
		user, _ := user.Current()
		return strings.Replace(path, "~/", user.HomeDir+string(os.PathSeparator), 1)
	} else {
		return path
	}
}

func (c *Config) AwsCredentials() *credentials.Credentials {
	return credentials.NewStaticCredentials(
		c.Aws.AccessKeyId,
		c.Aws.SecretAccessKey,
		"")
}
