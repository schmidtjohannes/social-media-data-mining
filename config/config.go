package config

import (
	"gopkg.in/yaml.v2"
)

type Filter struct {
	Keywords []string `yaml:"keywords"`
}

type Network struct {
	AccessToken string   `yaml:"access-token"`
	Groups      []string `yaml:"groups"`
}

type Configuration struct {
	Filter   Filter             `yaml:"filter"`
	Networks map[string]Network `yaml:"networks"`
}

type Parameters map[string]interface{}

func ParseConfiguration(yamlFile []byte) (*Configuration, error) {

	var c Configuration

	err := yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
