package table

import "gorm.io/gorm"

type Location struct {
	gorm.Model

	UserID    uint    `faker:"-"`
	Latitude  float32 `faker:"lat"`
	Longitude float32 `faker:"long"`
}

type Locations []Location
