package model

type Store struct {
	CustomerID  uint        `boil:"customer_id" json:"customer_id"`
	Name        string      `boil:"name" json:"name"`
	Description string      `boil:"description" json:"description"`
	Address     string      `boil:"address" json:"address"`
	Latitude    float64     `boil:"latitude" json:"latitude"`
	Longitude   float64     `boil:"longitude" json:"longitude"`
	Categories  *[]Category `boil:"categories" json:"categories,omitempty"`
	Clips       *[]Clip     `boil:"clips" json:"clips,omitempty"`
	Visits      *[]Visit    `boil:"visits" json:"visits,omitempty"`
}

type Stores []Store
