package table

import "gorm.io/gorm"

type Customer struct {
	gorm.Model

	Name              string `faker:"name"`
	Email             string `faker:"email"`
	Gender            string `faker:"word"`
	Age               int    `faker:"digit"`
	GoogleAccessToken string `faker:"uuid_hyphenated"`
	Store             Store  `gorm:"foreignKey:CustomerID"`
}

type Customers []Customer
