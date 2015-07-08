package model

type User struct {
	Id   int    `json:"id"`
	Name string `json:"username"`
}

type Users []User
