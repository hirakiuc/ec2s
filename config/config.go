package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"os/user"
	"strings"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

type Config struct {
	Aws  Aws
	Peco Peco
	Ssh  Ssh
}

type Aws struct {
	AccessKeyId     string `toml:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `toml:"AWS_SECRET_ACCESS_KEY"`
	Region          string `toml:"AWS_REGION"`
}

type Peco struct {
	Path string `toml:"path"`
}

type Ssh struct {
	Port          int            `toml:"port"`
	User          string         `toml:"user"`
	IdentityFiles []IdentityFile `toml:"identity_file"`
}

type IdentityFile struct {
	Name string `toml:"name"`
	Path string `toml:"path"`
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

func (ssh *Ssh) IdentityFileForName(name string) *string {
	for _, identityFile := range ssh.IdentityFiles {
		if identityFile.Name == name {
			return &identityFile.Path
		}
	}

	return nil
}
