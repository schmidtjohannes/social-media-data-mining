package miners

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"social-media-data-mining/config"
	"testing"
)

var fbNetwork = config.Network{
	AccessToken: "key1",
	Groups: []string{
		"group1",
		"group2",
	},
}

var body = `{
   "data":[
      {
         "message":"Contents of the Post",
         "id":"123456789123456789",
         "comments":{
            "data":[
               {
                  "message":"Contents of the Comment",
                  "from":{
                     "name":"John Doe",
                     "id":"123456789"
                  },
                  "likes":{
                     "data":[

                     ],
                     "summary":{
                        "total_count":14,
                        "can_like":true,
                        "has_liked":false
                     }
                  }
               }
            ]
         }
      }
   ]
}
`
var fbExpectedData = &FacebookGroupResponse{
	Items: []FacebookGroupItem{
		{Message: "Contents of the Post"},
	},
}

func TestFacebookMiner(t *testing.T) {

	fb := newFacebookMiner(fbNetwork)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, body)
	}))
	defer ts.Close()

	fb.url = ts.URL

	data, err := fb.QueryGroup()
	assert.Nil(t, err)
	assert.True(t, assert.ObjectsAreEqual(data, fbExpectedData))

	// http antwortet nicht
	// http ist != 200
	// http hat keine daten
	// message hat keine comments
	// message hat kein summary

}
