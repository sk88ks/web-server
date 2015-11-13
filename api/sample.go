package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sk88ks/web-server/entity"
)

// GetFriends is sample API
func GetFriends(c *gin.Context) {
	p := entity.Person{
		UserName: "Kokubo",
		Emails: []string{
			"test@awa.fm",
			"test@gmai.com",
		},
		Friends: []entity.Friend{
			{
				Fname: "Suzken",
			},
			{
				Fname: "Michinobu",
			},
		},
		HasContent: "having",
	}

	render(c, 200, "sample.templ", p)
}
