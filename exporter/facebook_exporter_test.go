package exporter

import (
	"fmt"
	"github.com/schmidtjohannes/social-media-data-mining/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestExportFacebookData(t *testing.T) {
	fbData := model.FacebookStatistics{
		GroupData: map[string]model.FacebookGroupStatistics{
			"group1": {
				Details: []model.FacebookStatisticDetail{
					{Post: "Post 1 group1", Likes: 37, Comments: 2},
				},
			},
			"group2": {
				Details: []model.FacebookStatisticDetail{
					{Post: "Post 1 group2 mo", Likes: 7, Comments: 1},
					{Post: "Post 2 group2 mo", Likes: 8, Comments: 2},
				},
			},
		},
	}
	expected := `Group-ID,Post,Likes,Comments
group1,Post 1 group1,37,2
group2,Post 1 group2 mo,7,1
group2,Post 2 group2 mo,8,2
`
	err := ExportFacebookData(fbData)
	assert.Nil(t, err)

	ts := getCurrentTimestamp()

	filename := fmt.Sprintf("stats-%s.csv", ts)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	assert.Equal(t, expected, str)
	os.Remove(filename)
}
