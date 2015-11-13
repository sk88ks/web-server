package main

import "github.com/gin-gonic/gin"

var itemQueryStrings = []QueryFunc{
	{Type: "findByPostId", Func: nil},
	{Type: "findByPostDateTimeGTE", Func: nil},
	{Type: "findByPostDateTimeLTE", Func: nil},
	{Type: "findByPostUserId", Func: nil},
	{Type: "findByPostItemId", Func: nil},
	{Type: "findByPostItemScoreGTE", Func: nil},
	{Type: "findByPostItemScoreLTE", Func: nil},
}

func SearchItem(c *gin.Context) {
	execFunc(c, itemQueryStrings)
}

func findByPostId(c *gin.Context, postID, limit string) {

}
