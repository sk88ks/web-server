package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sk88ks/web-server/entity"
)

// GetAlive return status
func GetAlive(c *gin.Context) {
	c.JSON(200, &entity.AliveRes{
		OK: true,
	})
}
