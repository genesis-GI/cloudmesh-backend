package main

import(
	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context){
	c.File("public/html/index.html")
}

func launcherDownloadHandler(c *gin.Context){
	c.File("public/html/launcherdownload.html")
}

func loginWebsiteHandler(c *gin.Context){
	c.File("public/html/login.html")
}

func regsiterWebsiteHandler(c *gin.Context){
	c.File("public/html/register.html")
}