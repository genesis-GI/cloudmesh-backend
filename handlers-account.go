package main

import (
	"fmt"
	"net/http"
	"time"

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
	var loginReq LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(loginReq)
	success, errMSG := login(loginReq)
	_=success

	if success {
		c.JSON(200, gin.H{
			"message": errMSG,
			"date": time.Now(),
		})

	}else{
		
		c.JSON(500, gin.H{
			"error": errMSG,
		})
	}
}