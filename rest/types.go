package main

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	password string `json:password`
}
