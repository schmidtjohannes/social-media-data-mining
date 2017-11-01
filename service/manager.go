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
	facebookMiner miners.FacebookManagerInterface
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
        fb, err := miners.NewFacebookManager(m.Config)
        if err != nil {
                return err
        }
	m.facebookMiner = fb
	return nil
}

func (m *DataMinerManager) execute() error {
	log.Print("Consuming social media networks")

	fbData, err := m.facebookMiner.QueryGroups()
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
