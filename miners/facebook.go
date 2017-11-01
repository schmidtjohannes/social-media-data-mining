package miners

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/schmidtjohannes/social-media-data-mining/model"
)

//API v2.10
var fbEndpoint = "https://graph.facebook.com/v2.10/"
var fbQuery = "/feed?fields=message,created_time,likes.limit(0).summary(true),comments.limit(10).summary(true){message,from,likes.limit(0).summary(true)}"
var fbMembersEndpoint = "members"

type FacebookMinerInterface interface {
	QueryGroup() model.FacebookGroupResponse
}

type FacebookMiner struct {
	accessToken string
	group       string
	httpClient  HttpClient
	url         string
}

type HttpClient interface {
	Get(string) (resp *http.Response, err error)
}

type FacebookHttpClient struct {
	httpclient *http.Client
}

func (c FacebookHttpClient) Get(url string) (resp *http.Response, err error) {
	return c.httpclient.Get(url)
}

func newFacebookClient() HttpClient {
	return &FacebookHttpClient{
		httpclient: &http.Client{Timeout: 10 * time.Second},
	}
}

func newFacebookMiner(fbGroup, fbAccessToken string) FacebookMiner {
	fbm := FacebookMiner{
		accessToken: fbAccessToken,
		group:       fbGroup,
		httpClient:  newFacebookClient(),
		url:         getUrl(fbGroup, fbAccessToken),
	}
	return fbm
}

func (fbm *FacebookMiner) QueryGroup() (*model.FacebookGroupResponse, error) {
	resp, err := fbm.httpClient.Get(fbm.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fgr := &model.FacebookGroupResponse{}
	err = json.NewDecoder(resp.Body).Decode(fgr)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, &HttpStatusNokError{Code: resp.StatusCode, Message: resp.Status}
	}
	if fgr.Items == nil || len(fgr.Items) == 0 {
		return nil, &EmptyResultError{Name: fbm.group}
	}
	return fgr, nil
}

func getUrl(group, accessToken string) string {
	return fmt.Sprintf("%s%s%s&access_token=%s", fbEndpoint, group, fbQuery, accessToken)
}
