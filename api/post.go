package main

import "github.com/gin-gonic/gin"

var postQueryStrings = []QueryFunc{}

// SearchPost is controller for API
func SearchPost(c *gin.Context) {
	execFunc(c, postQueryStrings)
}
