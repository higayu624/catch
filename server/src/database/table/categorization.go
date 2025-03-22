package table

import "gorm.io/gorm"

type Categorization struct {
	gorm.Model

	StoreId    uint
	CategoryId uint
}
