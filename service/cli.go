package service

import (
	"flag"
	"fmt"
	"log"
	"social-media-data-mining/logger"
)

func RunApp() {

	logger.SetupLogger()

	log.Print("Welcome to the data-miner")

	configFilePath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	log.Print(fmt.Sprintf("Config file path: %s", *configFilePath))

	dmManager, err := newDataMinerManager(*configFilePath)
        if err != nil {
                log.Fatal(fmt.Sprintf("ERROR: %s", err))
        }	
	err = dmManager.Execute()
	if err != nil {
		log.Fatal(fmt.Sprintf("ERROR: %s", err))
	}
}
