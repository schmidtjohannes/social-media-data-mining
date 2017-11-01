package analyzer

import (
	"fmt"
	"github.com/schmidtjohannes/social-media-data-mining/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFacebookAnalyzer(t *testing.T) {
	expectedResult := model.FacebookStatistics{
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
	fbInputGroup1 := &model.FacebookGroupResponse{
		Items: []model.FacebookGroupItem{
			{
				Message:     "Post 1 group1",
				Id:          "123456789123456789",
				CreatedTime: "2017-10-31T02:56:53+0000",
				Likes: model.Like{
					Summary: model.Summary{
						TotalCount: 37,
						CanLike:    true,
						HasLiked:   false,
					},
				},
				Comments: model.Comments{
					Data: []model.Comment{
						{
							Message: "Contents of the Comment",
							From: model.FacebookUser{
								Name: "John Doe",
								Id:   "123456789",
							},
							Likes: model.Like{
								Summary: model.Summary{
									TotalCount: 14,
									CanLike:    true,
									HasLiked:   false,
								},
							},
						},
						{
							Message: "Contents of the Comment two",
							From: model.FacebookUser{
								Name: "John Doe",
								Id:   "123456789",
							},
							Likes: model.Like{
								Summary: model.Summary{
									TotalCount: 2,
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
	fbInputGroup2 := &model.FacebookGroupResponse{
		Items: []model.FacebookGroupItem{
			{
				Message:     "Post 1 group2 more text here ",
				Id:          "123456789123456789",
				CreatedTime: "2017-10-31T02:56:53+0000",
				Likes: model.Like{
					Summary: model.Summary{
						TotalCount: 7,
						CanLike:    true,
						HasLiked:   false,
					},
				},
				Comments: model.Comments{
					Data: []model.Comment{
						{
							Message: "Contents of the Comment",
							From: model.FacebookUser{
								Name: "John Doe",
								Id:   "123456789",
							},
							Likes: model.Like{
								Summary: model.Summary{
									TotalCount: 14,
									CanLike:    true,
									HasLiked:   false,
								},
							},
						},
					},
				},
			},
			{
				Message:     "Post 2 group2 more text here ",
				Id:          "123456789123456789",
				CreatedTime: "2017-10-31T02:56:53+0000",
				Likes: model.Like{
					Summary: model.Summary{
						TotalCount: 8,
						CanLike:    true,
						HasLiked:   false,
					},
				},
				Comments: model.Comments{
					Data: []model.Comment{
						{
							Message: "Contents of the Comment post 2 group 2",
							From: model.FacebookUser{
								Name: "John Doe",
								Id:   "123456789",
							},
							Likes: model.Like{
								Summary: model.Summary{
									TotalCount: 14,
									CanLike:    true,
									HasLiked:   false,
								},
							},
						},
						{
							Message: "Contents of the Comment two post 2 group 2",
							From: model.FacebookUser{
								Name: "John Doe",
								Id:   "123456789",
							},
							Likes: model.Like{
								Summary: model.Summary{
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

	a := make(map[string]*model.FacebookGroupResponse)
	a["group1"] = fbInputGroup1
	a["group2"] = fbInputGroup2
	res := AnalyzeFacebookData(a)
	assert.True(t, assert.ObjectsAreEqual(res, expectedResult))
}
