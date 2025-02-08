package main

import (
	"os"
	"strings"
	"time"

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
	}else{
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

	download := r.Group("/download")
	{
		download.GET("/", func(c *gin.Context){
			c.JSON(200, gin.H{
				"!message":"The index route is unused, use rather: ",
				"Launcher download": "/download/launcher",
			})
		})

		download.GET("/launcher", func (c *gin.Context)  {
			launcherDownloadHandler(c)
		})
	}


	r.GET("/login", func(c *gin.Context){
		loginWebsiteHandler(c)
	})


	r.GET("/register", func(c *gin.Context){
		registerWebsiteHandler(c)
	})

	r.GET("/news", func(c *gin.Context){
		newsHandler(c)
	})


	r.POST("/login", func(c *gin.Context){
		POSTloginHandler(c)
	})

	r.POST("/register", func(c *gin.Context) {
		POSTregisterHandler(c)
	})

	r.GET("/ai", func(c *gin.Context){
		if gin.Mode() == gin.DebugMode  {
			c.File("public/html/ai.html")
		}else{
			c.String(503, "Service unavailable as the feature is not ready yet!")	
		}
	})

	r.GET("/connection/info", func(c *gin.Context){
		infoHandler(c)
	})

	r.GET("/versions/:email", func(c *gin.Context){
		if isDbEnabled {
			getVersions(c)
		}else {
			c.String(503, "Service only available with DB enabled!")
		}
	})


	r.POST("/motd/:motd", func(c *gin.Context){
		newMotd := c.Param("motd")
		currentTime := time.Now()
		success, msg := setMOTD(newMotd, currentTime)
		if !success{
			c.String(500, msg)
		}
		c.String(200, msg)
	})

	r.GET("/motd", func(c *gin.Context){
		current, lastupdate := getMOTD()
		c.JSON(200, gin.H{
			"message": current,
			"timestamp":lastupdate,
		})
	})

	r.NoRoute(func (c *gin.Context){
		noRouteHandler(c)
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