package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)



func main() { 
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	
	r.GET("/", func(c *gin.Context){
		indexHandler(c)
	})


	if(gin.Mode() == gin.DebugMode){
		fmt.Println("Running in debug mode")
	}
	fmt.Println("Server running on http://localhost:8088")
	r.Run(":8088")
}