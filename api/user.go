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
	{Type: "findByPostId", Func: findByPostID},
	{Type: "findByPostDateTimeGTE", Func: findByPostDateTimeGTE},
	{Type: "findByPostDateTimeLTE", Func: findByPostDateTimeLTE},
	{Type: "findByPostItemId", Func: findByPostItemID},
	{Type: "findByMaxPostItemScoreGTE", Func: findByMaxPostItemScoreGTE},
	{Type: "findByMinPostItemScoreLTE", Func: findByMinPostItemScoreLTE},
	{Type: "findByPostItemState", Func: findByPostItemState},
	{Type: "findByPostItemStateNotEQ", Func: findByPostItemStateNotEQ},
}

// SearchUser is for user API
func SearchUser(c *gin.Context) {
	execFunc(c, userQueryStrings)
}

func findByUserID(c *gin.Context, userID, limit string) {
	user, err := service.FindByUserID(userID)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.UserRes{
		Result: true,
		Data:   []entity.User{user},
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

func findByPostItemID(c *gin.Context, itemID, limit string) {
	users, err := service.FindByPostItemID(itemID, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.UserRes{
		Result: true,
		Data:   users,
	})
}

func findByMaxPostItemScoreGTE(c *gin.Context, score, limit string) {
	users, err := service.FindByMaxPostItemScoreGTE(score, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.UserRes{
		Result: true,
		Data:   users,
	})
}

func findByMinPostItemScoreLTE(c *gin.Context, score, limit string) {
	users, err := service.FindByMinPostItemScoreLTE(score, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.UserRes{
		Result: true,
		Data:   users,
	})
}

func findByPostItemState(c *gin.Context, state string, limit string) {
	users, err := service.FindByPostItemState(state, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.UserRes{
		Result: true,
		Data:   users,
	})
}

func findByPostItemStateNotEQ(c *gin.Context, state string, limit string) {
	users, err := service.FindByPostItemStateNotEQ(state, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.UserRes{
		Result: true,
		Data:   users,
	})
}
