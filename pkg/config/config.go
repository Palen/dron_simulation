package config

import (
	"github.com/jinzhu/configor"
)

// Configuration is stuff that can be configured externally per env variables or config file
type Configuration struct {
	SubscribersDir string `default:"./data/subscribers/"`
	CheckPointFile string `default:"./data/tube.csv"`
	Drone          struct {
		Perimeter float64 `default:"350.0"`
		Speed     float64 `default:"2.5"`
		MaxSize   int     `default:"10"`
	}
}

// Returns configuration files
func configFiles() []string {
	return []string{}
}

// Get returns the configuration extracted from env variables or config file.
func LoadConfig() (*Configuration, error) {
	conf := new(Configuration)
	err := configor.New(&configor.Config{ENVPrefix: "SIMULATION"}).
		Load(conf, configFiles()...)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
