package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func main() {
	configFile := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	fmt.Println(fmt.Sprintf("Config: %s", *configFile))

	fmt.Printf("hello, world\n")
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)
}
