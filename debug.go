package main

import(
	"os"
	"time"
	"fmt"
	"bufio"
	"github.com/gin-gonic/gin"
)

func debug(){
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
			debugMode = true
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
			debugMode = false
			gin.SetMode(gin.ReleaseMode)
		}
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout! Starting in release mode.")
		gin.SetMode(gin.ReleaseMode)
		debugMode = false
	}

}