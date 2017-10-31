package miners

import (
	"errors"
	"fmt"
	"github.com/schmidtjohannes/social-media-data-mining/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var fbNetwork = config.Network{
	AccessToken: "key1",
	Groups: []string{
		"group1",
		"group2",
	},
}

var emptyBody = `{
   "data":[]}`

var body = `{
   "data":[
      {
         "message":"Contents of the Post",
         "id":"123456789123456789",
	"created_time": "2017-10-31T02:56:53+0000",
      "likes": {
        "data": [
        ],
        "summary": {
          "total_count": 28,
          "can_like": true,
          "has_liked": false
        }
      },
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
		{
			Message:     "Contents of the Post",
			Id:          "123456789123456789",
			CreatedTime: "2017-10-31T02:56:53+0000",
			Likes: Like{
				Summary: Summary{
					TotalCount: 28,
					CanLike:    true,
					HasLiked:   false,
				},
			},
			Comments: Comments{
				Data: []Comment{
					{
						Message: "Contents of the Comment",
						From: FacebookUser{
							Name: "John Doe",
							Id:   "123456789",
						},
						Likes: Like{
							Summary: Summary{
								TotalCount: 14,
								CanLike:    true,
								HasLiked:   false,
							},
						},
					},
				},
			},
		},
	},
}

func TestFacebookMiner(t *testing.T) {

	tests := []struct {
		httpStatus   int
		httpResponse string
		success      bool
		mockClient   bool
	}{
		{http.StatusOK, body, true, false},
		{http.StatusInternalServerError, body, false, false},
		{http.StatusOK, emptyBody, false, false},
		{http.StatusOK, "", false, false},
		{http.StatusOK, "", false, true},
	}

	for _, v := range tests {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(v.httpStatus)
			fmt.Fprintln(w, v.httpResponse)
		}))
		defer ts.Close()

		fb := NewFacebookMiner("group1", "key1")
		fb.url = ts.URL

		if v.mockClient {
			fb.httpClient = &MockFacebookHttpClient{}
		}

		data, err := fb.QueryGroup()

		if v.success {
			assert.Nil(t, err)
			assert.True(t, assert.ObjectsAreEqual(data, fbExpectedData))
		} else {
			assert.NotNil(t, err)
			assert.Nil(t, data)
		}
	}
}

type MockFacebookHttpClient struct{}

func (m *MockFacebookHttpClient) Get(url string) (resp *http.Response, err error) {
	return nil, errors.New("fail")
}

/*
func TestFacebookMinerFailGet(t *testing.T) {

	fb := NewFacebookMiner("group1", "key1")
	fb.httpClient = &MockFacebookHttpClient{}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		return
	}))
	defer ts.Close()

	fb.url = ts.URL

	_, err := fb.QueryGroup()
	assert.NotNil(t, err)
}*/
