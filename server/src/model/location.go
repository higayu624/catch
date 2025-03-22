package model

type Location struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"user_id"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Locations []Location
