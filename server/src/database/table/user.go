package table

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name              string     `faker:"name"`
	Email             string     `faker:"email"`
	Gender            string     `faker:"word"`
	Age               int        `faker:"digit"`
	GoogleAccessToken string     `faker:"uuid_hyphenated"`
	Locations         []Location `gorm:"foreignKey:UserID"`
	Stores            []Store    `gorm:"many2many:visits;"`
}

type Users []User
