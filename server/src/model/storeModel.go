package model

type Store struct {
	CustomerID      uint              `boil:"customer_id" json:"customer_id"`
	Name            string            `boil:"name" json:"name"`
	Description     string            `boil:"description" json:"description"`
	Address         string            `boil:"address" json:"address"`
	Latitude        float64           `boil:"latitude" json:"latitude"`
	Longitude       float64           `boil:"longitude" json:"longitude"`
	Categorizations *[]Categorization `boil:"categorizations" json:"categorizations,omitempty"`
	Clips           *[]Clip           `boil:"clips" json:"clips,omitempty"`
	Users           *[]User           `boil:"users" json:"users,omitempty"`
}

type Stores []Store
