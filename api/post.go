package main

import "github.com/gin-gonic/gin"

var postQueryStrings = []QueryFunc{
	{Type: "findByPostId", Func: nil},
	{Type: "findByPostDateTimeGTE", Func: nil},
	{Type: "findByPostDateTimeLTE", Func: nil},
	{Type: "findByPostUserId", Func: nil},
	{Type: "findByPostItemId", Func: nil},
	{Type: "findByPostItemScoreGTE", Func: nil},
	{Type: "findByPostItemScoreLTE", Func: nil},
}

// SearchPost is controller for API
func SearchPost(c *gin.Context) {
	execFunc(c, postQueryStrings)
}
