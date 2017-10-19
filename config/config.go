package config

import(
	"gopkg.in/yaml.v2"
	"log"
)

type conf struct {
    Hits string `yaml:"hits"`
}

func (c *conf) getConf(yamlFile []byte) *conf {

    err := yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}
