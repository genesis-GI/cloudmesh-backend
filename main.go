package main

import (
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)


var useRemoteDB bool = true
var isDbEnabled bool = true
func main() {

	if len(os.Args) > 1 {
		input := os.Args[1]
		input = strings.ToLower(input)

		if input == "debug" {
			debug()
			gin.SetMode(gin.DebugMode)
			
		}else {
			color.Red("[✗ FAILURE] Invalid argument: %s", input)
			os.Exit(1)
		}
	}else
	{
		color.Cyan("[INFO]: No arguments provided")
		gin.SetMode(gin.ReleaseMode)
	}
	color.Cyan("[INFO]: Starting in %s mode", gin.Mode())	
	r := gin.Default()

	r.GET("/css/styles.css", func(c *gin.Context) {
		c.File("public/css/styles.css")
	})

	r.GET("/favicon.ico", func(c * gin.Context){
		c.String(200, "There is no icon currently!")
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
		if gin.Mode() == gin.ReleaseMode{
			c.String(503, "Service unavailable!")
		}else {
			regsiterWebsiteHandler(c)
		}
	})

	r.GET("/news", func(c *gin.Context){
		newshandler(c)
	})


	r.POST("/login", func(c *gin.Context){
		POSTloginHandler(c)
	})

	r.POST("/register/:email/:username/:password", func(c *gin.Context){
		POSTregisterHandler(c)
	})

	r.GET("/ai", func(c *gin.Context){
		if gin.Mode() == gin.DebugMode  {
			c.File("public/html/ai.html")
		}else{
			c.JSON(503, gin.H{
				"503":"Service unavailable!",
				"message": "This is under construction and will come soon!",
			})	
		}
	})

	
	r.NoRoute(func (c *gin.Context){
		errorHandler(c)
	})
	

	if isDbEnabled {
		err := initDB()
		if err != nil{
			panic(err)
		}
	}

	color.Magenta("[Environment]: %s", gin.Mode())
	color.Green("Server running on http://localhost:8088")
	r.Run(":8088")
}