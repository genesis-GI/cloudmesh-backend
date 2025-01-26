package main

import(
	"os"
	"fmt"
	"bufio"
	"github.com/gin-gonic/gin"

)

func debug(){
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
} 