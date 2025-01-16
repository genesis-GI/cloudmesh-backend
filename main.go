package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)



func main() { 
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.Static("/public/css", ".public/css")

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


	

	if(gin.Mode() == gin.DebugMode){
		fmt.Println("Running in debug mode")
	}
	fmt.Println("Server running on http://localhost:8088")
	r.Run(":8088")
}