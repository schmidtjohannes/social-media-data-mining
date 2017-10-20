package config

import (
	"gopkg.in/yaml.v2"
	valid "github.com/asaskevich/govalidator"
)

type conf struct {
	Hits string `yaml:"hits"`
}

func getConf(yamlFile []byte) (*conf, error) {

	var c conf

	err := yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, err
	}

	if _, err := valid.ValidateStruct(c); err != nil {
		return nil, err
	}

	return &c, nil
}
