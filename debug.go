package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"strings"
)

func debug(){
	gin.SetMode(gin.DebugMode)

	color.Magenta("[Environment] Debug Mode enabled!")
	fmt.Println("Do you want to disable database? (y/n)")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if input == "y\n" || input == "y\r\n" {
		color.Cyan("[INFO] Database is disabled...")
		isDbEnabled = false
	}else{
		fmt.Println("Do you want to use local database? (y/n)")
		input, _ = reader.ReadString('\n')
		if input == "y\n" || input == "y\r\n" {
			color.Cyan("[INFO] Using local database...")
			useRemoteDB = false
		}
	}
} 

func getParameters() {
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
		gin.SetMode(gin.ReleaseMode)
	}
}