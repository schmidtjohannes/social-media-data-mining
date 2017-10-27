package miners

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"social-media-data-mining/config"
	"time"
)

//API v2.10
var fbEndpoint = "https://graph.facebook.com/v2.10/"
var fbQuery = "/feed?fields=message,comments.limit(10).summary(true){message,from,likes.limit(0).summary(true)}"

type FacebookGroupResponse struct {
	Items []FacebookGroupItem `json:"data"`
}
type FacebookGroupItem struct {
	Message  string   `json:"message"`
	Id       string   `json:"id"`
	Comments Comments `json:"comments"`
}

type Comments struct {
	Data []Comment `json:"data"`
}

type Comment struct {
	Message string       `json:"message"`
	From    FacebookUser `json:"from"`
	Likes   Like         `json:"likes"`
}

type Like struct {
	Summary Summary `json:"summary"`
}

type Summary struct {
	TotalCount int64 `json:"total_count"`
	CanLike    bool  `json:"can_like"`
	HasLiked   bool  `json:"has_liked"`
}

type FacebookUser struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type FacebookMinerInterface interface {
	QueryGroup() FacebookGroupResponse
}

type FacebookMiner struct {
	accessToken string
	groups      []string
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

func newFacebookMiner(config config.Network) FacebookMiner {
	fbm := FacebookMiner{
		accessToken: config.AccessToken,
		groups:      config.Groups,
		httpClient:  newFacebookClient(),
		url:         getUrl(config.Groups[0], config.AccessToken),
	}
	return fbm
}

func (fbm *FacebookMiner) QueryGroup() (*FacebookGroupResponse, error) {
	log.Print(fbm.url)
	resp, err := fbm.httpClient.Get(fbm.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fgr := &FacebookGroupResponse{}
	err = json.NewDecoder(resp.Body).Decode(fgr)
	if err != nil {
		return nil, err
	}
	return fgr, nil
}

func getUrl(group, accessToken string) string {
	return url.QueryEscape(fmt.Sprintf("%s%s%s&access_token=%s", fbEndpoint, group, fbQuery, accessToken))
}
