package miners

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//API v2.10
var fbEndpoint = "https://graph.facebook.com/v2.10/"
var fbQuery = "/feed?fields=message,created_time,likes.limit(0).summary(true),comments.limit(10).summary(true){message,from,likes.limit(0).summary(true)}"

type FacebookGroupResponse struct {
	Items []FacebookGroupItem `json:"data"`
}
type FacebookGroupItem struct {
	Message     string   `json:"message"`
	CreatedTime string   `json:"created_time"`
	Id          string   `json:"id"`
	Comments    Comments `json:"comments"`
	Likes       Like     `json:"likes"`
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

func NewFacebookMiner(fbGroup, fbAccessToken string) FacebookMiner {
	fbm := FacebookMiner{
		accessToken: fbAccessToken,
		group:       fbGroup,
		httpClient:  newFacebookClient(),
		url:         getUrl(fbGroup, fbAccessToken),
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
