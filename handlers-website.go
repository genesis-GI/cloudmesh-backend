package main

import(
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
	c.File("public/html/error.html")
}

func infoHandler(c *gin.Context){
	remote :=c.RemoteIP()
	client := c.ClientIP()
	method := c.Request.Method
	uri := c.Request.RequestURI
	protocol := c.Request.Proto

	c.String(200, "Remote: "+remote+"\n\nClient: "+client+"\n\nMethod: "+method+"\n\nProtocol: "+protocol+"\n\nURI: "+uri)
}