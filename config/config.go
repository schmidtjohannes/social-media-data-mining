package config

import (
	"errors"
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

type Parser struct{}
type ParserInterface interface {
	ParseConfiguration([]byte) (*Configuration, error)
}

func (p *Parser) ParseConfiguration(yamlFile []byte) (*Configuration, error) {

	var c Configuration

	err := yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, err
	}
	if c.Networks == nil || len(c.Networks) == 0 {
		return nil, errors.New("no network found")
	}
	return &c, nil
}
