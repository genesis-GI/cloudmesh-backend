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

}

type registerRequest struct{

}