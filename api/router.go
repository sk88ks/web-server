package main

import "github.com/gin-gonic/gin"

func router(r *gin.Engine) {

	r.GET("/alive", GetAlive)
}
