package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"github.com/gin-gonic/gin"
)


var useRemoteDB bool = true

func main() {

	var isDbEnabled bool = true
	

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
			gin.SetMode(gin.DebugMode)
			fmt.Println("Debug mode enabled!")
			fmt.Println("Do you want to disable database? (y/n)")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			if input == "y\n" || input == "y\r\n" {
				fmt.Println("Database is disabled.")
				isDbEnabled = false
			}else{
				fmt.Println("Do you want to use remote database? (y/n)")
				input, _ = reader.ReadString('\n')
				if input == "n\n" || input == "n\r\n" {
					fmt.Println("Using local database.")
					useRemoteDB = false
				}
			}
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


	
	r.NoRoute(func (c *gin.Context){
		indexHandler(c)
	})
	

	if isDbEnabled {
		err := initDB()
		if err != nil{
			panic(err)
		}
	}

	fmt.Println("Environment:", gin.Mode())
	fmt.Println("Server running on http://localhost:8088")
	r.Run(":8088")
}