package table

import "gorm.io/gorm"

type Visit struct {
	gorm.Model

	UserId  uint `json:"user_id"`
	StoreId uint `json:"store_id"`
}

type Visits []Visit
