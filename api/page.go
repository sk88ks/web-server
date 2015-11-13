package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sk88ks/web-server/entity"
	"github.com/sk88ks/web-server/service"
)

func UserPage(c *gin.Context) {
	userID := c.Param("userid")
	user, err := service.FindByUserID(userID)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	data := entity.UserPage{
		Self: user,
	}

	render(c, 200, "user.templ", data)
}
