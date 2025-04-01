package model

type Categorization struct {
	StoreID    uint `boil:"store_id" json:"store_id"`
	CategoryID uint `boil:"category_id" json:"category_id"`
}
