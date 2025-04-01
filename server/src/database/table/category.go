package table

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	Name            string           `faker:"restaurant_category"`
	Categorizations []Categorization `gorm:"foreignKey:CategoryID"`
}
