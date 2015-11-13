package datastore

import (
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
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(
			&user.UserID,
			&user.UserNo,
			&user.UserPublicScore,
			&user.UserFriends,
			&user.UserImage,
		)
		if err != nil {
			return nil, err
		}
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
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := entity.Item{}
		err = rows.Scan(
			&item.IttemID,
			&item.ItemNo,
			&item.ItemSupplier,
			&item.ItemSoldQuantity,
			&item.ItemSalePrice,
			&item.ItemTags,
			&item.ItemImage,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
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
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		post := entity.Post{}
		err = rows.Scan(
			&post.PostID,
			&post.PostDateTime,
			&post.PostUserID,
			&post.PostItemID,
			&post.PostItemScore,
			&post.PostItemState,
			&post.PossLikeUsers,
			&post.PostTags,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
