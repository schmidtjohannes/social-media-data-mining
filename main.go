package main

import (
	"flag"
	"fmt"
)

func main() {
	configFile := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	fmt.Println(fmt.Sprintf("Config: %s", *configFile))
}
