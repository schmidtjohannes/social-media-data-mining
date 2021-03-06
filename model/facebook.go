package model

type FacebookStatistics struct {
	GroupData map[string]FacebookGroupStatistics
}

type FacebookGroupStatistics struct {
	Details []FacebookStatisticDetail
}

type FacebookStatisticDetail struct {
	Post     string
	Likes    int64
	Comments int
}

type FacebookGroupResponse struct {
	Items []FacebookGroupItem `json:"data"`
}
type FacebookGroupItem struct {
	Message     string   `json:"message"`
	Story       string   `json:"story"`
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
