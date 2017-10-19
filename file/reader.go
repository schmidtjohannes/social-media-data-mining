package file

import (
	"io/ioutil"
)

func readFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
