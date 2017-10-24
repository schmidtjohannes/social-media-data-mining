package service

import (
        "log"
        "social-media-data-mining/config"
        "social-media-data-mining/file"
)

type DataMinerManager struct {
	Config *config.Configuration
}

func newDataMinerManager(configFilePath string) (*DataMinerManager, error) {
        log.Print("Reading config file")
        cfgFile, err := file.ReadFile(configFilePath)
        if err != nil {
                return nil, err
        }
        log.Print("Parsing config file")
        cfgStruct, err := config.ParseConfiguration(cfgFile)
        if err != nil {
                return nil, err
        }
	dm := DataMinerManager{ Config : cfgStruct }
	return &dm, nil 
}

func (m *DataMinerManager) Execute() error {
        log.Print("Consuming social media networks")
        //miner
        log.Print("Analyzing data")
        //analyzer
        log.Print("Exporting data")
        //exporter
	return nil
}
