package model

type Clip struct {
	ID          uint   `boil:"id" json:"id"`
	StoreID     uint   `boil:"store_id" json:"store_id"`
	Name        string `boil:"name" json:"name"`
	S3Token     string `boil:"s3_token" json:"s3_token"`
	Description string `boil:"description" json:"description"`
}

type Clips []Clip
