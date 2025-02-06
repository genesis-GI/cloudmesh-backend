package main

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context){
	c.File("public/html/index.html")
}

func launcherDownloadHandler(c *gin.Context){
	c.File("public/html/launcher.html")
}

func loginWebsiteHandler(c *gin.Context){
	c.File("public/html/login.html")
}

func registerWebsiteHandler(c *gin.Context){
	c.File("public/html/register.html")
}
func newsHandler(c *gin.Context){
	c.File("public/html/news.html")
}

func noRouteHandler(c *gin.Context){
	htmlContent, err := os.ReadFile("public/html/error.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading the page")
		return
	}
	c.Data(http.StatusNotFound, "text/html; charset=utf-8", htmlContent)
}

func infoHandler(c *gin.Context){
	remote :=c.RemoteIP()
	client := c.ClientIP()
	method := c.Request.Method
	uri := c.Request.RequestURI
	protocol := c.Request.Proto

	c.JSON(200, gin.H{
		"Remote": remote,
		"Client": client,
		"Method": method,
		"Protocol": protocol,
		"URI": uri,
	})

	//c.String(200, "Remote: "+remote+"\n\nClient: "+client+"\n\nMethod: "+method+"\n\nProtocol: "+protocol+"\n\nURI: "+uri)
}