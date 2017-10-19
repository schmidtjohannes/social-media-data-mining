package config

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

var data = `
hits: Easy!
`

var corruptData = `
nofield: Easy!
`
func TestConfig(t *testing.T){
	c := conf{}
	c.getConf([]byte(data))
	assert.NotNil(t, c)
	assert.Equal(t, "Easy!", c.Hits)
}

func TestBrokenConfig(t *testing.T){
        c := conf{}
        c.getConf([]byte(corruptData))
        assert.Nil(t, c)
}
