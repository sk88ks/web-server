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

	q := "SELECT * FROM post4 WHERE id = '" + postID + "'"
	posts, err := datastore.PostQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	itemIDMap := make(map[string]bool, len(posts))
	for i := range posts {
		if itemIDMap[posts[i].PostItemID] {
			continue
		}
		itemIDMap[posts[i].PostItemID] = true
	}

	var itemIDs string
	for id := range itemIDMap {
		itemIDs += `'` + id + `',`
	}

	q = "SELECT * FROM item WHERE id IN(" + itemIDs + ") ORDER BY userNo LIMIT " + limit
	items, err := datastore.ItemQueryWithCache(q)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, nil
	}

	return items, nil
}

//func FindByPostDateTimeGTE(unixtime, limit string) ([]entity.Item, error) {
//	if limit == "" {
//		limit = "100"
//	}
//
//	q := "SELECT * FROM post WHERE id = '" + postID + "'"
//	posts, err := datastore.PostQueryWithCache(q)
//	if err != nil {
//		return nil, err
//	}
//
//	itemIDMap := make(map[string]bool, len(posts))
//	for i := range posts {
//		if itemIDMap[posts[i].PostItemID] {
//			continue
//		}
//		itemIDMap[posts[i].PostItemID] = true
//	}
//
//	var itemIDs string
//	for id := range itemIDMap {
//		itemIDs += `'` + id + `',`
//	}
//
//	q = "SELECT * FROM item WHERE id IN(" + itemIDs + ") ORDER BY userNo LIMIT " + limit
//	items, err := datastore.ItemQueryWithCache(q)
//	if err != nil {
//		return nil, err
//	}
//
//	if len(items) == 0 {
//		return nil, nil
//	}
//
//	return items, nil
//
//}