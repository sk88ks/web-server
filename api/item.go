package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sk88ks/web-server/entity"
	"github.com/sk88ks/web-server/service"
)

var itemQueryStrings = []QueryFunc{
	{Type: "findByItemId", Func: findItemByItemID},
	{Type: "findByItemSupplier", Func: findItemBySupplier},
	{Type: "findByItemSoldQuantityGTE", Func: findItemByItemSoldQuantityGTE},
	{Type: "findByItemSoldQuantityLTE", Func: findItemByItemSoldQuantityLTE},
	{Type: "findByItemSalePriceGTE", Func: findItemByItemSalePriceGTE},
	{Type: "findByItemSalePriceLTE", Func: findItemByItemSalePriceLTE},
	{Type: "findByItemTagsIncludeAll", Func: dummyFunc},
	{Type: "findByItemTagsIncludeAny", Func: dummyFunc},
	{Type: "findByPostId", Func: findItemByPostID},
	{Type: "findByPostDateTimeGTE", Func: nil},
	{Type: "findByPostDateTimeLTE", Func: nil},
}

func SearchItem(c *gin.Context) {
	execFunc(c, itemQueryStrings)
}

func findItemByItemID(c *gin.Context, itemID, limit string) {
	item, err := service.FindItemByItemID(itemID)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.ItemRes{
		Result: true,
		Data:   []entity.Item{item},
	})
}

func findItemBySupplier(c *gin.Context, supplier, limit string) {
	items, err := service.FindItemBySupplier(supplier, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.ItemRes{
		Result: true,
		Data:   items,
	})
}

func findItemByItemSoldQuantityGTE(c *gin.Context, quantity, limit string) {
	items, err := service.FindItemByItemSoldQuantityGTE(quantity, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.ItemRes{
		Result: true,
		Data:   items,
	})
}

func findItemByItemSoldQuantityLTE(c *gin.Context, quantity, limit string) {
	items, err := service.FindItemByItemSoldQuantityLTE(quantity, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.ItemRes{
		Result: true,
		Data:   items,
	})
}

func findItemByItemSalePriceGTE(c *gin.Context, price, limit string) {
	items, err := service.FindItemByItemSalePriceGTE(price, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.ItemRes{
		Result: true,
		Data:   items,
	})
}

func findItemByItemSalePriceLTE(c *gin.Context, price, limit string) {
	items, err := service.FindItemByItemSalePriceLTE(price, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.ItemRes{
		Result: true,
		Data:   items,
	})
}

func findItemByPostID(c *gin.Context, postID, limit string) {
	items, err := service.FindItemByPostID(postID, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.ItemRes{
		Result: true,
		Data:   items,
	})
}
