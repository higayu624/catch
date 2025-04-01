package table

import "gorm.io/gorm"

type Categorization struct {
	gorm.Model

	StoreID    uint `boil:"store_id" json:"store_id"`
	CategoryID uint `boil:"category_id" json:"category_id"`
}
