package config

import (
	"email/configParser"
	"path"
)

const (
	emailConfigFileName = "../conf/email.toml"
)

var EmailConfig configParser.EmailConfig

func LoadConfig(configPath string) error {
	err := configParser.ReadConfig(path.Join(configPath, emailConfigFileName), &EmailConfig)
	if err != nil {
		return err
	}

	return nil
}
