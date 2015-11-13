package datastore

import (
	"fmt"
	"strings"

	"github.com/sk88ks/web-server/entity"
	"github.com/sk88ks/web-server/env"
)

// UserQueryWithCache is qury for user
func UserQueryWithCache(query string) ([]entity.User, error) {
	// TODO use cache

	conn, err := env.GetMySQLConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	users := []entity.User{}
	fmt.Printf("User Query: %s\n", query)
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := entity.User{}
		var friends string
		err = rows.Scan(
			&user.UserID,
			&user.UserNo,
			&user.UserPublicScore,
			&friends,
			&user.UserImage,
		)
		if err != nil {
			return nil, err
		}
		user.UserFriends = strings.Split(friends, ",")
		users = append(users, user)
	}

	return users, nil
}

// ItemQueryWithCache is qury for user
func ItemQueryWithCache(query string) ([]entity.Item, error) {
	// TODO use cache

	conn, err := env.GetMySQLConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	items := []entity.Item{}
	fmt.Printf("Item Query: %s\n", query)
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := entity.Item{}
		var tags string
		err = rows.Scan(
			&item.IttemID,
			&item.ItemNo,
			&item.ItemSupplier,
			&item.ItemSoldQuantity,
			&item.ItemSalePrice,
			&tags,
			&item.ItemImage,
		)
		if err != nil {
			return nil, err
		}
		item.ItemTags = strings.Split(tags, ",")
		items = append(items, item)
	}

	return items, nil
}

func PostQueryForUserID(query string) ([]string, error) {
	conn, err := env.GetMySQLConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	fmt.Printf("Posts Query: %s\n", query)
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	userIDs := []string{}
	for rows.Next() {
		var userID string
		err = rows.Scan(
			&userID,
		)
		if err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}

	return userIDs, nil
}

func PostQueryForItemID(query string) ([]string, error) {
	conn, err := env.GetMySQLConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	fmt.Printf("Posts Query: %s\n", query)
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	itemIDs := []string{}
	for rows.Next() {
		var itemID string
		err = rows.Scan(
			&itemID,
		)
		if err != nil {
			return nil, err
		}
		itemIDs = append(itemIDs, itemID)
	}

	return itemIDs, nil
}

// PostQueryWithCache is qury for user
func PostQueryWithCache(query string) ([]entity.Post, error) {
	// TODO use cache

	conn, err := env.GetMySQLConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	posts := []entity.Post{}
	fmt.Printf("Posts Query: %s\n", query)
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		post := entity.Post{}
		var lusers string
		var tags string
		err = rows.Scan(
			&post.PostID,
			&post.PostDateTime,
			&post.PostUserID,
			&post.PostItemID,
			&post.PostItemScore,
			&post.PostItemState,
			&lusers,
			&tags,
		)
		if err != nil {
			return nil, err
		}
		post.PostTags = strings.Split(tags, ",")
		post.PostLikeUsers = strings.Split(lusers, ",")
		posts = append(posts, post)
	}

	return posts, nil
}
