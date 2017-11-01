package config

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

var corruptData = `
filter=1
`

var configYaml = `
filter:
  keywords:
    - escuela
    - alumnos
    - padres
networks:
  facebook:
    access-token: key1
    groups:
      - group1
      - group2
`
var configNoNetwork = `
filter:
  keywords:
    - escuela
    - alumnos
    - padres
`

var configStruct = Configuration{
	Filter: Filter{
		Keywords: []string{
			"escuela",
			"alumnos",
			"padres",
		},
	},
	Networks: map[string]Network{
		"facebook": {
			AccessToken: "key1",
			Groups: []string{
				"group1",
				"group2",
			},
		},
	},
}

func TestBrokenConfig(t *testing.T) {
	p := new(Parser)
	c, err := p.ParseConfiguration([]byte(corruptData))
	assert.NotNil(t, err)
	assert.Nil(t, c)
}

func TestParseSimple(t *testing.T) {
	p := new(Parser)
	config, err := p.ParseConfiguration([]byte(configYaml))
	assert.Nil(t, err)
	assert.Equal(t, config, &configStruct)
}

func TestMarshalRoundtrip(t *testing.T) {
	configBytes, err := yaml.Marshal(configStruct)
	assert.Nil(t, err)
	p := new(Parser)
	config, err := p.ParseConfiguration(configBytes)

	assert.Nil(t, err)
	assert.True(t, assert.ObjectsAreEqual(config, &configStruct))
}

func TestConfigAtLeastOneNetwork(t *testing.T){
        p := new(Parser)
        config, err := p.ParseConfiguration([]byte(configNoNetwork))
	assert.NotNil(t, err)
	assert.Nil(t, config)
}
