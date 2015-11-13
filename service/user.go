package service

import (
	"github.com/sk88ks/web-server/datastore"
	"github.com/sk88ks/web-server/entity"
	"github.com/sk88ks/web-server/env"
)

func FindByUserID(userID string) (user entity.User, err error) {
	if userID == "" {
		return
	}

	q := "SELECT * FROM user WHERE id = '" + userID + "'"

	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return
	}

	if len(users) == 0 {
		return
	}

	return users[0], nil
}

func FindByUserPublicScoreGTE(num, limit string) ([]entity.User, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT * FROM user WHERE userPublicScore >= " + num + " ORDER BY userNo LIMIT " + limit
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

	q := "SELECT * FROM user WHERE userPublicScore <= " + num + " ORDER BY userNo LIMIT " + limit
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

	q := "SELECT * FROM user WHERE friendsNum >= " + num + " ORDER BY userNo LIMIT " + limit
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

	q := "SELECT * FROM user WHERE friendsNum <= " + num + " ORDER BY userNo LIMIT " + limit
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

	q := "SELECT * FROM post WHERE postId = '" + postID + "'"
	posts, err := datastore.PostQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, nil
	}

	q = "SELECT * FROM user WHERE id = '" + posts[0].PostUserID + "' ORDER BY userNo LIMIT " + limit
	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

func FindByPostDateTimeGTE(unixtime, limit string) ([]entity.User, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT * FROM post WHERE postDatetime >= " + unixtime
	posts, err := datastore.PostQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, nil
	}

	userIDMap := make(map[string]bool, len(posts))
	for i := range posts {
		if userIDMap[posts[i].PostUserID] {
			continue
		}
		userIDMap[posts[i].PostUserID] = true
	}

	var userIDs string
	for id := range userIDMap {
		userIDs += `'` + id + `',`
	}

	q = "SELECT * FROM user WHERE id IN(" + userIDs + ") ORDER BY userNo LIMIT " + limit
	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

func FindByPostDateTimeLTE(unixtime, limit string) ([]entity.User, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT * FROM post WHERE postDatetime <= " + unixtime
	posts, err := datastore.PostQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, nil
	}

	userIDMap := make(map[string]bool, len(posts))
	for i := range posts {
		if userIDMap[posts[i].PostUserID] {
			continue
		}
		userIDMap[posts[i].PostUserID] = true
	}

	var userIDs string
	for id := range userIDMap {
		userIDs += `'` + id + `',`
	}

	q = "SELECT * FROM user WHERE id IN(" + userIDs + ") ORDER BY userNo LIMIT " + limit
	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

func FindByPostItemID(itemID, limit string) ([]entity.User, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT * FROM post WHERE postItemId = " + itemID
	posts, err := datastore.PostQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, nil
	}

	userIDMap := make(map[string]bool, len(posts))
	for i := range posts {
		if userIDMap[posts[i].PostUserID] {
			continue
		}
		userIDMap[posts[i].PostUserID] = true
	}

	var userIDs string
	for id := range userIDMap {
		userIDs += `'` + id + `',`
	}

	q = "SELECT * FROM user WHERE id IN(" + userIDs + ") ORDER BY userNo LIMIT " + limit
	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

func FindByMaxPostItemScoreGTE(score, limit string) ([]entity.User, error) {
	if limit == "" {
		limit = "100"
	}

	conn, err := env.GetMySQLConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	q := `SELECT * FROM post WHERE postItemScore >= ` + score
	posts, err := datastore.PostQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, nil
	}

	userIDMap := make(map[string]bool, len(posts))
	for i := range posts {
		if userIDMap[posts[i].PostUserID] {
			continue
		}
		userIDMap[posts[i].PostUserID] = true
	}

	var userIDs string
	for id := range userIDMap {
		userIDs += `'` + id + `',`
	}

	q = "SELECT * FROM user WHERE id IN(" + userIDs + ") ORDER BY userNo LIMIT " + limit
	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

func FindByMinPostItemScoreLTE(score, limit string) ([]entity.User, error) {
	if limit == "" {
		limit = "100"
	}

	conn, err := env.GetMySQLConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	q := `SELECT * FROM post WHERE postItemScore <= ` + score
	posts, err := datastore.PostQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, nil
	}

	userIDMap := make(map[string]bool, len(posts))
	for i := range posts {
		if userIDMap[posts[i].PostUserID] {
			continue
		}
		userIDMap[posts[i].PostUserID] = true
	}

	var userIDs string
	for id := range userIDMap {
		userIDs += `'` + id + `',`
	}

	q = "SELECT * FROM user WHERE id IN(" + userIDs + ") ORDER BY userNo LIMIT " + limit
	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

func FindByPostItemState(state, limit string) ([]entity.User, error) {
	if limit == "" {
		limit = "100"
	}

	conn, err := env.GetMySQLConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	q := `SELECT * FROM post WHERE postItemState = ` + state
	posts, err := datastore.PostQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, nil
	}

	userIDMap := make(map[string]bool, len(posts))
	for i := range posts {
		if userIDMap[posts[i].PostUserID] {
			continue
		}
		userIDMap[posts[i].PostUserID] = true
	}

	var userIDs string
	for id := range userIDMap {
		userIDs += `'` + id + `',`
	}

	q = "SELECT * FROM user WHERE id IN(" + userIDs + ") ORDER BY userNo LIMIT " + limit
	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

func FindByPostItemStateNotEQ(state, limit string) ([]entity.User, error) {
	if limit == "" {
		limit = "100"
	}

	conn, err := env.GetMySQLConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	q := `SELECT * FROM post WHERE postItemState <> ` + state
	posts, err := datastore.PostQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, nil
	}

	userIDMap := make(map[string]bool, len(posts))
	for i := range posts {
		if userIDMap[posts[i].PostUserID] {
			continue
		}
		userIDMap[posts[i].PostUserID] = true
	}

	var userIDs string
	for id := range userIDMap {
		userIDs += `'` + id + `',`
	}

	q = "SELECT * FROM user WHERE id IN(" + userIDs + ") ORDER BY userNo LIMIT " + limit
	users, err := datastore.UserQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}
