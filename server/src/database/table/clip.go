package table

import "gorm.io/gorm"

type Clip struct {
	gorm.Model

	StoreID     uint
	Name        string `faker:"name"`
	S3Token     string `faker:"uuid_hyphenated"`
	Description string `faker:"paragraph"`
}
