package miners

import (
	"fmt"
	"github.com/schmidtjohannes/social-media-data-mining/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var configStruct = &config.Configuration{
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

func TestFacebookNoNetwork(t *testing.T) {
	cfg := &config.Configuration{
		Filter: config.Filter{
			Keywords: []string{
				"escuela",
				"alumnos",
				"padres",
			},
		}}
	fbm, err := NewFacebookManager(cfg)
	assert.NotNil(t, err)
	assert.Nil(t, fbm)
}

func TestFacebookNoGroups(t *testing.T) {
	cfg := &config.Configuration{
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
			},
		},
	}
	fbm, err := NewFacebookManager(cfg)
	assert.NotNil(t, err)
	assert.Nil(t, fbm)
}

func TestFacebookNoAccesskey(t *testing.T) {
	cfg := &config.Configuration{
		Filter: config.Filter{
			Keywords: []string{
				"escuela",
				"alumnos",
				"padres",
			},
		},
		Networks: map[string]config.Network{
			"facebook": {
				Groups: []string{
					"group1",
					"group2",
				},
			},
		},
	}
	fbm, err := NewFacebookManager(cfg)
	assert.NotNil(t, err)
	assert.Nil(t, fbm)
}

func TestFacebookHapppyPath(t *testing.T) {
	fbm, err := NewFacebookManager(configStruct)
	assert.Nil(t, err)
	assert.NotNil(t, fbm)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, body)
	}))
	defer ts.Close()

	fbEndpoint = fmt.Sprintf("%s/", ts.URL)
	fbQuery = "/test"

	fbRes, err := fbm.QueryGroups()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(fbRes))
}
func TestFacebookError(t *testing.T) {
	fbm, err := NewFacebookManager(configStruct)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "fail")
	}))
	defer ts.Close()

	fbEndpoint = fmt.Sprintf("%s/", ts.URL)
	fbQuery = "/test"

	fbRes, err := fbm.QueryGroups()
	assert.NotNil(t, err)
	assert.Nil(t, fbRes)
}
