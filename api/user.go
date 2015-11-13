package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sk88ks/web-server/entity"
	"github.com/sk88ks/web-server/service"
)

var userQueryStrings = []QueryFunc{
	{Type: "findByUserId", Func: findUserByUserID},
	{Type: "findByUserPublicScoreGTE", Func: findUserByUserPublicScoreGTE},
	{Type: "findByUserPublicScoreLTE", Func: findUserByUserPublicScoreLTE},
	{Type: "findByUserFriendsNumberGTE", Func: findUserByUserFriendsNumberGTE},
	{Type: "findByUserFriendsNumberLTE", Func: findUserByUserFriendsNumberLTE},
	{Type: "findByUserFriendsIncludeUserIds", Func: dummyFunc},
	{Type: "findByUserFriendsNotIncludeUserIds", Func: dummyFunc},
	{Type: "findByPostId", Func: findUserByPostID},
	{Type: "findByPostDateTimeGTE", Func: findUserByPostDateTimeGTE},
	{Type: "findByPostDateTimeLTE", Func: findUserByPostDateTimeLTE},
	{Type: "findByPostItemId", Func: findUserByPostItemID},
	{Type: "findByMaxPostItemScoreGTE", Func: findUserByMaxPostItemScoreGTE},
	{Type: "findByMinPostItemScoreLTE", Func: findUserByMinPostItemScoreLTE},
	{Type: "findByPostItemState", Func: findUserByPostItemState},
	{Type: "findByPostItemStateNotEQ", Func: findUserByPostItemStateNotEQ},
}

// SearchUser is for user API
func SearchUser(c *gin.Context) {
	execFunc(c, userQueryStrings)
}

func findUserByUserID(c *gin.Context, userID, limit string) {
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

func findUserByUserPublicScoreGTE(c *gin.Context, num, limit string) {
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

func findUserByUserPublicScoreLTE(c *gin.Context, num, limit string) {
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

func findUserByUserFriendsNumberGTE(c *gin.Context, num, limit string) {
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

func findUserByUserFriendsNumberLTE(c *gin.Context, num, limit string) {
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

func findUserByPostID(c *gin.Context, postID, limit string) {
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

func findUserByPostDateTimeGTE(c *gin.Context, unixtime, limit string) {
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

func findUserByPostDateTimeLTE(c *gin.Context, unixtime, limit string) {
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

func findUserByPostItemID(c *gin.Context, itemID, limit string) {
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

func findUserByMaxPostItemScoreGTE(c *gin.Context, score, limit string) {
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

func findUserByMinPostItemScoreLTE(c *gin.Context, score, limit string) {
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

func findUserByPostItemState(c *gin.Context, state string, limit string) {
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

func findUserByPostItemStateNotEQ(c *gin.Context, state string, limit string) {
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
