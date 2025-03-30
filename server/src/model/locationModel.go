package model

type Location struct {
	ID        uint    `boil:"id" json:"id"`
	UserID    uint    `boil:"user_id" json:"user_id"`
	Latitude  float32 `boil:"latitude" json:"latitude"`
	Longitude float32 `boil:"longitude" json:"longitude"`
}

type Locations []Location
