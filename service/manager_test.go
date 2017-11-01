package service

import (
	"errors"
	"github.com/schmidtjohannes/social-media-data-mining/config"
	"github.com/schmidtjohannes/social-media-data-mining/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

type FacebookManagerMock struct {
	data []*model.FacebookGroupResponse
	error error
}

func (fmm *FacebookManagerMock) QueryGroups() ([]*model.FacebookGroupResponse, error) {
	return fmm.data, fmm.error
}

type FileReaderMock struct {
	mockReadFile func(string) ([]byte, error)
}

func (f *FileReaderMock) ReadFile(path string) ([]byte, error) {
	if f.mockReadFile != nil {
		return f.mockReadFile(path)
	}
	return ioutil.ReadFile("../file/test-config.yaml")
}

type ParserMock struct {
	mockParseConfiguration func([]byte) (*config.Configuration, error)
}

func (p *ParserMock) ParseConfiguration(yamlFile []byte) (*config.Configuration, error) {
	if p.mockParseConfiguration != nil {
		return p.mockParseConfiguration(yamlFile)
	}
	cfg := config.Configuration{
		Filter: config.Filter{
			Keywords: []string{
				"escuela",
				"alumnos",
				"padres",
			},
		},
		Networks: map[string]config.Network{
			"facebook": {
				AccessToken: "key1",
				Groups: []string{
					"group1",
					"group2",
				},
			},
		},
	}

	return &cfg, nil
}

func TestCreateManager(t *testing.T) {
	dm := newDataMinerManager()
	assert.NotNil(t, dm)
}

func TestInit(t *testing.T) {
	dm := newDataMinerManager()
	dm.setFileReader(new(FileReaderMock))
	dm.setParser(new(ParserMock))
	err := dm.init("")
	assert.Nil(t, err)

	frMock := &FileReaderMock{
		mockReadFile: func(string) ([]byte, error) {
			return nil, errors.New("failed")
		},
	}
	dm.setFileReader(frMock)
	err = dm.init("")
	assert.NotNil(t, err)

	dm.setFileReader(new(FileReaderMock))
	pMock := &ParserMock{
		mockParseConfiguration: func([]byte) (*config.Configuration, error) {
			return nil, errors.New("failed")
		},
	}
	dm.setParser(pMock)
	err = dm.init("")
	assert.NotNil(t, err)
}

func TestExecuteManager(t *testing.T) {
	dm := new(DataMinerManager)
        dm.setFileReader(new(FileReaderMock))
        dm.setParser(new(ParserMock))
	dm.init("")
	mockResponse := &model.FacebookGroupResponse{
		Items: []model.FacebookGroupItem{
			{ Message:     "Contents of the Post", },
		},
	}
	a := []*model.FacebookGroupResponse {mockResponse}

	dm.facebookMiner = &FacebookManagerMock{ data : a, error : nil }
	res := dm.execute()
	assert.Nil(t, res)
	
	dm.facebookMiner = &FacebookManagerMock{ data : nil, error : errors.New("fail") }
        res = dm.execute()
        assert.NotNil(t, res)
}
