package model

type Categorization struct {
	StoreId    uint `json:"store_id"`
	CategoryId uint `json:"category_id"`
}

type categorizations []Categorization
