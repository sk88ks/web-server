package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sk88ks/web-server/entity"
	"github.com/sk88ks/web-server/service"
)

var itemQueryStrings = []QueryFunc{
	{Type: "findByItemId", Func: findByItemID},
	{Type: "findByItemSupplier", Func: findBySupplier},
	{Type: "findByItemSoldQuantityGTE", Func: findByItemSoldQuantityGTE},
	{Type: "findByItemSoldQuantityLTE", Func: findByItemSoldQuantityLTE},
	{Type: "findByItemSalePriceGTE", Func: findByItemSalePriceGTE},
	{Type: "findByItemSalePriceLTE", Func: findByItemSalePriceLTE},
	{Type: "findByItemTagsIncludeAll", Func: nil},
	{Type: "findByItemTagsIncludeAny", Func: nil},
	{Type: "findByPostId", Func: nil},
}

func SearchItem(c *gin.Context) {
	execFunc(c, itemQueryStrings)
}

func findByItemID(c *gin.Context, itemID, limit string) {
	item, err := service.FindByItemID(itemID)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.ItemRes{
		Result: true,
		Data:   []entity.Item{item},
	})
}

func findBySupplier(c *gin.Context, supplier, limit string) {
	items, err := service.FindBySupplier(supplier, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.ItemRes{
		Result: true,
		Data:   items,
	})
}

func findByItemSoldQuantityGTE(c *gin.Context, quantity, limit string) {
	items, err := service.FindByItemSoldQuantityGTE(quantity, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.ItemRes{
		Result: true,
		Data:   items,
	})
}

func findByItemSoldQuantityLTE(c *gin.Context, quantity, limit string) {
	items, err := service.FindByItemSoldQuantityLTE(quantity, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.ItemRes{
		Result: true,
		Data:   items,
	})
}

func findByItemSalePriceGTE(c *gin.Context, price, limit string) {
	items, err := service.FindByItemSalePriceGTE(price, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.ItemRes{
		Result: true,
		Data:   items,
	})
}

func findByItemSalePriceLTE(c *gin.Context, price, limit string) {
	items, err := service.FindByItemSalePriceLTE(price, limit)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, entity.ItemRes{
		Result: true,
		Data:   items,
	})
}

func findByPostID(c *gin.Context, postID, limit string) {

}
