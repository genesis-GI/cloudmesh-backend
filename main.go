package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Press 'd' and Enter within 5 seconds to enable debug mode...")
	debugModeCh := make(chan bool)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if input == "d\n" || input == "d\r\n" { 
			debugModeCh <- true
		} else {
			debugModeCh <- false
		}
	}()
	select {
	case enableDebug := <-debugModeCh:
		if enableDebug {
			fmt.Println("Debug mode enabled!")
			gin.SetMode(gin.DebugMode)
		} else {
			fmt.Println("Starting in release mode.")
			gin.SetMode(gin.ReleaseMode)
		}
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout! Starting in release mode.")
		gin.SetMode(gin.ReleaseMode)
	}



	r := gin.Default()

	
	r.GET("/css/styles.css", func(c *gin.Context) {
		c.File("public/css/styles.css")
	})

	r.GET("/favicon.ico", func(c * gin.Context){
		c.String(200, "Comming soon!")
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

	r.GET("/news", func(c *gin.Context){
		newshandler(c)
	})


	r.POST("/login/:email/:password", func(c *gin.Context){
		POSTloginHandler(c)

	})

	r.POST("/register/:email/:username/:password", func(c *gin.Context){
			POSTregisterHandler(c)
	})


	

	if(gin.Mode() == gin.DebugMode){
		fmt.Println("Running in debug mode...")
		fmt.Println("...Database is disabled")
	}else{
		err := initDB()
		if err != nil{
			panic(err)
		}
	}
	 
	fmt.Println("Server running on http://localhost:8088")
	r.Run(":8088")
}