package analyzer

import (
	"github.com/schmidtjohannes/social-media-data-mining/model"
)

func AnalyzeFacebookData(fbData map[string]*model.FacebookGroupResponse) model.FacebookStatistics {
	fbGroupStatistics := make(map[string]model.FacebookGroupStatistics)
	for key, groupResponse := range fbData {
		fbGroupStatistics[key] = model.FacebookGroupStatistics{
			Details: extractPosts(groupResponse.Items),
		}
	}
	return model.FacebookStatistics{GroupData: fbGroupStatistics}
}

func extractPosts(posts []model.FacebookGroupItem) []model.FacebookStatisticDetail {
	var fbStatsDetailSlice []model.FacebookStatisticDetail
	for postIdx := range posts {
		fbStatsDetailSlice = append(fbStatsDetailSlice, extractDetails(posts[postIdx]))
	}
	return fbStatsDetailSlice
}

func extractDetails(post model.FacebookGroupItem) model.FacebookStatisticDetail {
	fbStatsDetail := model.FacebookStatisticDetail{
		Post:     short(getPostContent(post), 24),
		Likes:    post.Likes.Summary.TotalCount,
		Comments: len(post.Comments.Data),
	}
	return fbStatsDetail
}

func getPostContent(post model.FacebookGroupItem) string {
	if post.Message == "" {
		return post.Story
	}
	return post.Message
}

func short(s string, i int) string {
	runes := []rune(s)
	if len(runes) > i {
		return string(runes[:i])
	}
	return s
}
