package config

import (
	"github.com/jinzhu/configor"
)

// Configuration is stuff that can be configured externally per env variables or config file
type Configuration struct {
	Database struct {
		Dialect  string `default:"mongo"`
		User     string `default:""`
		Password string `default:""`
		Host     string `default:"db"`
		DBName   string `default:"incomedb"`
		Port     uint   `default:"27017"`
	}
	FTP struct {
		Host     string `default:"ftp"`
		Port     uint   `default:"21"`
		User     string `default:"anonymous"`
		Password string `default:""`
		Timer    int    `default:"300"`
	}
}

// Returns configuration files
func configFiles() []string {
	return []string{}
}

// Get returns the configuration extracted from env variables or config file.
func LoadConfig() (*Configuration, error) {
	conf := new(Configuration)
	err := configor.New(&configor.Config{ENVPrefix: "CONFIG"}).
		Load(conf, configFiles()...)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
