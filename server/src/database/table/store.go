package table

import "gorm.io/gorm"

type Store struct {
	gorm.Model

	CustomerID      uint             `faker:"-"`
	Name            string           `faker:"name"`
	Description     string           `faker:"paragraph"`
	Address         string           `faker:"address"`
	Latitude        float32          `faker:"lat"`
	Longitude       float32          `faker:"long"`
	Categorizations []Categorization `gorm:"foreignKey:StoreID"`
	Clips           []Clip           `gorm:"foreignKey:StoreID"`
	Users           []User           `gorm:"many2many:visits;"`
}

type Stores []Store
