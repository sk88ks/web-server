package service

import (
	"github.com/sk88ks/web-server/datastore"
	"github.com/sk88ks/web-server/entity"
)

func FindByItemID(itemID string) (item entity.Item, err error) {
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

func FindBySupplier(supplier, limit string) ([]entity.Item, error) {
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

func FindByItemSoldQuantityGTE(quantity, limit string) ([]entity.Item, error) {
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

func FindByItemSoldQuantityLTE(quantity, limit string) ([]entity.Item, error) {
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

func FindByItemSalePriceGTE(price, limit string) ([]entity.Item, error) {
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

func FindByItemSalePriceLTE(price, limit string) ([]entity.Item, error) {
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

//func findByPostID(postID, limit string) ([]entity.Item, error) {
//
//}
