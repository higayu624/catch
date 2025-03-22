package model

type User struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Gender            string `json:"gender"`
	Age               string `json:"age"`
	GoogleAccessToken string `json:"google_access_token"`
	Locations         []Location
	Visits            []Visit
}

type Users []User
