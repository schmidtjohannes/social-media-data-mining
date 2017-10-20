package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var data = `
hits: Easy!
`

var corruptData = `
hits=1
`

func TestConfig(t *testing.T) {
	c, err := getConf([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "Easy!", c.Hits)
}

func TestBrokenConfig(t *testing.T) {
	c, err := getConf([]byte(corruptData))
	assert.NotNil(t, err)
	assert.Nil(t, c)
}
