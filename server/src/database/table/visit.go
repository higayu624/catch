package table

import "gorm.io/gorm"

type Visit struct {
	gorm.Model

	UserID  uint `boil:"user_id" json:"user_id"`
	StoreID uint `boil:"store_id" json:"store_id"`
}
