package model

type Clip struct {
	ID          uint   `json:"id"`
	StoreID     uint   `json:"store_id"`
	Name        string `json:"name"`
	S3Token     string `json:"s3_token"`
	Description string `json:"description"`
}

type Clips []Clip
