package main

import (
"github.com/gin-gonic/gin"

)

func POSTregisterHandler(c *gin.Context){
	email := c.Param("email")
	username := c.Param("username")
	password := c.Param("password")
	success, msg := register(email, username, password)

	if !success{
		c.String(500,msg)
	}else{
		c.String(201, msg)
	}
}

func POSTloginHandler(c *gin.Context){
	email := c.Param("email")
	password := c.Param("password")

	loginReq := LoginRequest {
		Email: email,
		Password: password,
	}

	success, errMSG := login(loginReq)

	if !success {
		c.String(500, errMSG)
	}else{
		c.String(200, errMSG)
	}
}