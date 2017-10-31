package service

import (
	"github.com/schmidtjohannes/social-media-data-mining/config"
	"github.com/schmidtjohannes/social-media-data-mining/file"
	"github.com/schmidtjohannes/social-media-data-mining/miners"
	"log"
)

type DataMinerManager struct {
	fileReader file.FileReaderInterface
	parser     config.ParserInterface
	Config     *config.Configuration
}

type DataMinerManagerInterface interface {
	execute() error
	init(string) error
	setFileReader(file.FileReaderInterface)
	setParser(config.ParserInterface)
}

func newDataMinerManager() DataMinerManagerInterface {
	dm := new(DataMinerManager)
	dm.fileReader = new(file.FileReader)
	dm.parser = new(config.Parser)
	return dm
}

func (m *DataMinerManager) setFileReader(fr file.FileReaderInterface) {
	m.fileReader = fr
}

func (m *DataMinerManager) setParser(p config.ParserInterface) {
	m.parser = p
}

func (m *DataMinerManager) init(configFilePath string) error {
	log.Print("Reading config file")
	cfgFile, err := m.fileReader.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	log.Print("Parsing config file")
	cfgStruct, err := m.parser.ParseConfiguration(cfgFile)
	if err != nil {
		return err
	}
	m.Config = cfgStruct
	return nil
}

func (m *DataMinerManager) execute() error {
	log.Print("Consuming social media networks")

	//todo: check if network exist
	fb := miners.NewFacebookMiner(m.Config.Networks["facebook"])

	fbData, err := fb.QueryGroup()
	if err != nil {
		return err
	}
	log.Print(fbData)
	log.Print("Analyzing data")
	//analyzer
	log.Print("Exporting data")
	//exporter
	return nil
}
