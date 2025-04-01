package model

type CreateCustomer struct {
	Name              string    `boil:"name" json:"name"`
	Email             string    `boil:"email" json:"email"`
	Gender            string    `boil:"gender" json:"gender"`
	Age               int64     `boil:"age" json:"age"`
	GoogleAccessToken string    `boil:"google_access_token" json:"google_access_token"`
	Store             *GetStore `boil:"store" json:"store,omitempty"`
}

type GetStore struct {
	CustomerID  uint       `boil:"customer_id" json:"customer_id"`
	Name        string     `boil:"name" json:"name"`
	Description string     `boil:"description" json:"description"`
	Address     string     `boil:"address" json:"address"`
	Latitude    float64    `boil:"latitude" json:"latitude"`
	Longitude   float64    `boil:"longitude" json:"longitude"`
	Categories  []Category `boil:"categorizations" json:"categorizations,omitempty"`
}
