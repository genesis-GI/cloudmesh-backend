package main

import "time"

type Account struct {
	Username  string    `bson:"username" json:"username"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"-"`
	Admin     bool      `bson:"admin" json:"admin"`
	Wave      int       `bson:"wave" json:"wave"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
