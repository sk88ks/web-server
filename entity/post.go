package entity

// Post is data for post
type Post struct {
	PostID        string   `json:"postId"`
	PostDateTime  int64    `json:"postDateTime"`
	PostUserID    string   `json:"postUserId"`
	PostItemID    string   `json:"pistItemId"`
	PostItemScore int      `json:"postItemScore"`
	PostItemState string   `json:"postItemState"`
	PostLikeUsers []string `json:"postLikeUsers"`
	PostTags      []string `json:"postTags"`
}

// PostRes is response for post
type PostRes struct {
	Result bool   `json:"result"`
	Data   []Post `json:"data"`
}
