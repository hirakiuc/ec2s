package config

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

// Config describe configuration of ec2s.
type Config struct {
	Aws    Aws
	Peco   Peco
	SSH    SSH
	Common Common
}

// Aws define config about aws.
type Aws struct {
	AccessKeyID     string `toml:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `toml:"AWS_SECRET_ACCESS_KEY"`
	Region          string `toml:"AWS_REGION"`
}

// Peco define config about peco
type Peco struct {
	Path string `toml:"path"`
}

// SSH define config about ssh.
type SSH struct {
	Port          int            `toml:"port"`
	User          string         `toml:"user"`
	IdentityFiles []IdentityFile `toml:"identity_file"`
}

// IdentityFile define config about identity file of ssh.
type IdentityFile struct {
	Name string `toml:"name"`
	Path string `toml:"path"`
}

// Common define configs about common.
type Common struct {
	ColorizedOutput bool `toml:"colorized_output"`
}

var config Config

// GetConfig create Config instance.
func GetConfig() *Config {
	return &config
}

// LoadConfig load toml config in the path.
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
	}

	return path
}

// AwsCredentials return credentials.Credentials instance.
func (c *Config) AwsCredentials() *credentials.Credentials {
	return credentials.NewStaticCredentials(
		c.Aws.AccessKeyID,
		c.Aws.SecretAccessKey,
		"")
}

// IdentityFileForName return path of the IdentityFile.
func (ssh *SSH) IdentityFileForName(name string) (*string, error) {
	for _, identityFile := range ssh.IdentityFiles {
		if identityFile.Name == name {
			return &identityFile.Path, nil
		}
	}

	return nil, fmt.Errorf("Can't find IdentityFile for %s", name)
}
