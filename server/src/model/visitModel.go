package model

type Visit struct {
	UserId  uint `boil:"user_id" json:"user_id"`
	StoreId uint `boil:"store_id" json:"store_id"`
}

type Visits []Visit
