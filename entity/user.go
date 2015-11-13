package entity

// AliveRes is response for alive API
type AliveRes struct {
	OK bool `json:"ok"`
}

// User is data res
type User struct {
	UserID          string   `json:"userId"`
	UserNo          int      `json:"userNo"`
	UserPublicScore int      `json:"userPublicScore"`
	UserFriends     []string `json:"userFriends"`
	UserImage       string   `json:"userImage"`
}

// UserRes is response
type UserRes struct {
	Result bool   `json:"result"`
	Data   []User `json:"data"`
}
