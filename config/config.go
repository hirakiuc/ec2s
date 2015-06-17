package config

import (
	"github.com/BurntSushi/toml"
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
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
