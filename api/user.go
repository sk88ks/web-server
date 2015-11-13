package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sk88ks/web-server/entity"
	"github.com/sk88ks/web-server/service"
)

var userQueryStrings = []QueryFunc{
	{Type: "findByUserId", Func: findByUserID},
	{Type: "findByUserPublicScoreGTE", Func: findByUserPublicScoreGTE},
	{Type: "findByUserPublicScoreLTE", Func: findByUserPublicScoreLTE},
	{Type: "findByUserFriendsNumberGTE", Func: findByUserFriendsNumberGTE},
	{Type: "findByUserFriendsNumberLTE", Func: findByUserFriendsNumberLTE},
	{Type: "findByUserFriendsIncludeUserIds", Func: nil},
	{Type: "findByUserFriendsNotIncludeUserIds", Func: nil},
	{Type: "findByPostId", Func: nil},
	{Type: "findByPostDateTimeGTE", Func: nil},
	{Type: "findByPostDateTimeLTE", Func: nil},
	{Type: "findByPostItemId", Func: nil},
	{Type: "findByMaxPostItemScoreGTE", Func: nil},
	{Type: "findByMinPostItemScoreLTE", Func: nil},
	{Type: "findByPostItemState", Func: nil},
	{Type: "findByPostItemStateNotEQ", Func: nil},
}

// SearchUser is for user API
func SearchUser(c *gin.Context) {
	execFunc(c, userQueryStrings)
}

func findByUserID(c *gin.Context, userID, limit string) {
	users, err := service.FindByUserID(userID, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	if len(users) == 0 {
		c.JSON(200, entity.UserRes{
			Result: true,
		})
		return
	}

	c.JSON(200, entity.UserRes{
		Result: true,
		Data:   []entity.User{users[0]},
	})
}

func findByUserPublicScoreGTE(c *gin.Context, num, limit string) {
	users, err := service.FindByUserPublicScoreGTE(num, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.UserRes{
		Result: true,
		Data:   users,
	})
}

func findByUserPublicScoreLTE(c *gin.Context, num, limit string) {
	users, err := service.FindByUserPublicScoreLTE(num, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.UserRes{
		Result: true,
		Data:   users,
	})
}

func findByUserFriendsNumberGTE(c *gin.Context, num, limit string) {
	users, err := service.FindByUserFriendsNumberGTE(num, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.UserRes{
		Result: true,
		Data:   users,
	})
}

func findByUserFriendsNumberLTE(c *gin.Context, num, limit string) {
	users, err := service.FindByUserFriendsNumberLTE(num, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.UserRes{
		Result: true,
		Data:   users,
	})
}

func findByPostID(c *gin.Context, postID, limit string) {
	users, err := service.FindByPostID(postID, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.UserRes{
		Result: true,
		Data:   users,
	})
}

func findByPostDateTimeGTE(c *gin.Context, unixtime, limit string) {
	users, err := service.FindByPostDateTimeGTE(unixtime, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.UserRes{
		Result: true,
		Data:   users,
	})
}

func findByPostDateTimeLTE(c *gin.Context, unixtime, limit string) {
	users, err := service.FindByPostDateTimeLTE(unixtime, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.UserRes{
		Result: true,
		Data:   users,
	})
}
