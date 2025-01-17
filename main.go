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


		// ******* TEMPORARY TEST *******
		hashedPassword, err := HashPassword("test")
		if err != nil {
			fmt.Println("Error:",err)
		}
		// ******* TEMPORARY TEST *******

		
		fmt.Println("Search db for email:",email)
		// Now, compare the password of the found document with the checkpassword function  (replace hashedpassword with the thing from the document)

		match := CheckPasswordHash(password, hashedPassword)
		

		if match{
			c.Status(200)
		}else{
			c.Status(500)
		}
		
	})

	r.POST("/register/:email/:username/:password", func(c *gin.Context){
		email := c.Param("email")
		username := c.Param("username")
		password := c.Param("password")
		registerReq := registerRequest{email: email, username: username, password: password}
		

		fmt.Println(registerReq)
	})




	if(gin.Mode() == gin.DebugMode){
		fmt.Println("Running in debug mode")
	}
	fmt.Println("Server running on http://localhost:8088")
	r.Run(":8088")
}