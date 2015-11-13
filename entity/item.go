package entity

// Item is item data
type Item struct {
	IttemID          string   `json:"itemId"`
	ItemNo           int      `json:"itemNo"`
	ItemSupplier     string   `json:"itemSupplier"`
	ItemSoldQuantity int      `json:"itemSoldQuantitu"`
	ItemSalePrice    int      `json:"itemSalePrice"`
	ItemTags         []string `json:"itemTags"`
	ItemImage        string   `json:"itemImage"`
}

// ItemRes is response
type ItemRes struct {
	Result bool   `json:"result"`
	Data   []Item `json:"data"`
}
