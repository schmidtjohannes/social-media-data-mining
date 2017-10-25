package file

import (
	"io/ioutil"
)

type FileReader struct {
}

type FileReaderInterface interface {
	ReadFile(string) ([]byte, error)
}

func (f *FileReader) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
