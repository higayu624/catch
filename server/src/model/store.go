package model

import (
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type Store struct {
	CustomerID      uint              `json:"customer_id"`
	Name            null.String       `json:"name"`
	Description     null.String       `json:"paragraph"`
	Address         null.String       `json:"address"`
	Latitude        types.NullDecimal `json:"latitude"`
	Longitude       types.NullDecimal `json:"Longitude"`
	Categorizations []Categorization
	Clips           []Clip
	Visits          []Visit
}

type Stores []Store
