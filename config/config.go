package config

import (
	"tool/configParser"
	"path"
)

const (
	configFileName = "./conf/conf.toml"
)

var C configParser.Config

func LoadConfig(configPath string) error {
	err := configParser.ReadConfig(path.Join(configPath, configFileName), &C)
	if err != nil {
		return err
	}

	return nil
}
