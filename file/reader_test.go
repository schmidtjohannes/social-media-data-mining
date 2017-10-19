package file

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReader(t *testing.T) {
	file, err := readFile("test-config.yaml")
	assert.Nil(t, err)
	assert.NotNil(t, file)
}

func TestReaderNoFile(t *testing.T) {
        file, err := readFile("test-broken-config.yaml")
        assert.NotNil(t, err)
        assert.Nil(t, file)
}
