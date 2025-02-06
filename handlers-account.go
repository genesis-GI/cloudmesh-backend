package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


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
func POSTregisterHandler(c *gin.Context){
	var req struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	success, message := register(req.Email, req.Username, req.Password)
	if success {
		c.JSON(200, gin.H{"message": message})
	} else {
		c.JSON(400, gin.H{"error": message})
	}
}