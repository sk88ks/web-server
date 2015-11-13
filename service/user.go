package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sk88ks/web-server/datastore"
	"github.com/sk88ks/web-server/entity"
)

func FindByUserID(userID, limit string) ([]entity.User, error) {
	if userID == "" {
		return nil, nil
	}

	if limit == "" {
		limit = "100"
	}

	q := "SELECT * FROM user WHERE userId = " + userID + "ORDER BY userNo LIMIT " + limit

	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func FindByUserPublicScoreGTE(num, limit string) ([]entity.User, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT * FROM user WHERE userPublicScore >= " + num + "ORDER BY userNo LIMIT " + limit
	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func FindByUserPublicScoreLTE(num, limit string) ([]entity.User, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT * FROM user WHERE userPublicScore <= " + num + "ORDER BY userNo LIMIT " + limit
	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func FindByUserFriendsNumberGTE(num, limit string) ([]entity.User, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT * FROM user WHERE friendsNum >= " + num + "ORDER BY userNo LIMIT " + limit
	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func FindByUserFriendsNumberLTE(num, limit string) ([]entity.User, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT * FROM user WHERE friendsNum <= " + num + "ORDER BY userNo LIMIT " + limit
	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func FindByPostID(postID, limit string) ([]entity.User, error) {
	if limit == "" {
		limit = "100"
	}

	if postID == "" {
		return nil, nil
	}

	q := "SELECT * FROM post WHERE postId = " + postID
	posts, err := datastore.PostQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, nil
	}

	q = "SELECT * FROM user WHERE userId =" + posts[0].PostUserID + "ORDER BY userNo LIMIT " + limit
	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

func findByPostDateTimeGTE(c *gin.Context, unixtime, limit string) {
	// 	if limit == "" {
	//		limit = "100"
	//	}
	//
	//	q := "SELECT * FROM post WHERE postDatetime >= " + unixtime
	//	posts, err := datastore.PostQueryWithCache(q)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	if len(posts) == 0 {
	//		return nil, nil
	//	}
	//
	//	var userIDs string
	//	for i := range posts {
	//		userIDs += posts[i].PostUserID + ","
	//	}
	//
	//
	//
}

func findByPostDateTimeLTE(c *gin.Context, unixtime, limit string) {
}