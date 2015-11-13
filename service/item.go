package service

import (
	"github.com/sk88ks/web-server/datastore"
	"github.com/sk88ks/web-server/entity"
)

func FindItemByItemID(itemID string) (item entity.Item, err error) {
	if itemID == "" {
		return
	}

	q := "SELECT * FROM item WHERE id = '" + itemID + "'"

	items, err := datastore.ItemQueryWithCache(q)
	if err != nil {
		return
	}

	if len(items) == 0 {
		return
	}

	return items[0], nil
}

func FindItemBySupplier(supplier, limit string) ([]entity.Item, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT * FROM item WHERE itemSupplier = '" + supplier + "' ORDER BY itemNo LIMIT " + limit
	items, err := datastore.ItemQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func FindItemByItemSoldQuantityGTE(quantity, limit string) ([]entity.Item, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT * FROM item WHERE itemSoldQuantity >= " + quantity + " ORDER BY itemNo LIMIT " + limit
	items, err := datastore.ItemQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func FindItemByItemSoldQuantityLTE(quantity, limit string) ([]entity.Item, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT * FROM item WHERE itemSoldQuantity <= " + quantity + " ORDER BY itemNo LIMIT " + limit
	items, err := datastore.ItemQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func FindItemByItemSalePriceGTE(price, limit string) ([]entity.Item, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT * FROM item WHERE itemSalePrice >= " + price + " ORDER BY itemNo LIMIT " + limit
	items, err := datastore.ItemQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func FindItemByItemSalePriceLTE(price, limit string) ([]entity.Item, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT * FROM item WHERE itemSalePrice <= " + price + " ORDER BY itemNo LIMIT " + limit
	items, err := datastore.ItemQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func FindItemByPostID(postID, limit string) ([]entity.Item, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT postItemId FROM post4 WHERE id = '" + postID + "'"
	posts, err := datastore.PostQueryForItemID(q)
	if err != nil {
		return nil, err
	}

	itemIDs := make([]byte, 11*len(posts))
	for i := range posts {
		itemIDs = append(itemIDs, "'"...)
		itemIDs = append(itemIDs, posts[i]...)
		itemIDs = append(itemIDs, "',"...)
	}

	q = "SELECT * FROM item WHERE id IN(" + string(itemIDs) + ") ORDER BY userNo LIMIT " + limit
	items, err := datastore.ItemQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, nil
	}

	return items, nil
}

func FindItemByPostDateTimeGTE(unixtime, limit string) ([]entity.Item, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT postItemId FROM post4 WHERE postDateTime >= " + unixtime
	posts, err := datastore.PostQueryForItemID(q)
	if err != nil {
		return nil, err
	}

	itemIDs := make([]byte, 11*len(posts))
	for i := range posts {
		itemIDs = append(itemIDs, "'"...)
		itemIDs = append(itemIDs, posts[i]...)
		itemIDs = append(itemIDs, "',"...)
	}

	q = "SELECT * FROM item WHERE id IN(" + string(itemIDs) + ") ORDER BY userNo LIMIT " + limit
	items, err := datastore.ItemQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, nil
	}

	return items, nil

}

func FindItemByPostDateTimeLTE(unixtime, limit string) ([]entity.Item, error) {
	if limit == "" {
		limit = "100"
	}

	q := "SELECT postItemId FROM post4 WHERE postDateTime <= " + unixtime
	posts, err := datastore.PostQueryForItemID(q)
	if err != nil {
		return nil, err
	}

	itemIDs := make([]byte, 11*len(posts))
	for i := range posts {
		itemIDs = append(itemIDs, "'"...)
		itemIDs = append(itemIDs, posts[i]...)
		itemIDs = append(itemIDs, "',"...)
	}

	q = "SELECT * FROM item WHERE id IN(" + string(itemIDs) + ") ORDER BY userNo LIMIT " + limit
	items, err := datastore.ItemQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, nil
	}

	return items, nil

}
