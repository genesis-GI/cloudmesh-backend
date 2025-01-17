package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)



func main() { 
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.GET("/css/styles.css", func(c *gin.Context) {
		c.File("public/css/styles.css")
	})


	
	r.GET("/", func(c *gin.Context){
		indexHandler(c)
	})

	r.GET("/download/launcher", func(c* gin.Context){
		launcherDownloadHandler(c)
	})

	r.GET("/login", func(c *gin.Context){
		loginWebsiteHandler(c)
	})


	r.GET("/register", func(c *gin.Context){
		regsiterWebsiteHandler(c)
	})


	r.POST("/login/:email/:password", func(c *gin.Context){
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

	})

	r.POST("/register/:email/:username/:password", func(c *gin.Context){
		email := c.Param("email")
		username := c.Param("username")
		password := c.Param("password")
		success, msg := register(email, username, password)

		if !success{
			c.String(500,msg)
		}else{
			c.String(201, msg)
		}
	})




	if(gin.Mode() == gin.DebugMode){
		fmt.Println("Running in debug mode")
	}
	initDB()
	fmt.Println("Server running on http://localhost:8088")
	r.Run(":8088")
}