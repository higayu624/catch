package model

type Visit struct {
	UserId  uint `json:"user_id"`
	StoreId uint `json:"store_id"`
}

type Visits []Visit
