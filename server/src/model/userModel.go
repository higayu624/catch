package model

type User struct {
	ID                string     `boil:"id" json:"id"`
	Name              string     `boil:"name" json:"name"`
	Email             string     `boil:"email" json:"email"`
	Gender            string     `boil:"gender" json:"gender"`
	Age               string     `boil:"age" json:"age"`
	GoogleAccessToken string     `boil:"google_access_token" json:"google_access_token"`
	Locations         []Location `boil:"locations" json:"locations"`
	Stores            []Store    `boil:"stores" json:"stores"`
}

type Users []User
