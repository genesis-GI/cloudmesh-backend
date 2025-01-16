package main


type account struct{
	username string
	email string
	password string
	admin bool
	wave int
	ingame map[string]interface{}
	playerLocation  map[string]interface{}
}

type loginRequest struct{
	email string
	password string
}

type registerRequest struct{
	email string
	username string
	password string
}