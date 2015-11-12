package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sk88ks/web-server/body"
)

// GetAlive return status
func GetAlive(c *gin.Context) {
	c.JSON(200, &body.AliveRes{
		OK: true,
	})
}
