package main

import (

	"os"

	"time"
	"github.com/fatih/color"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)


var useRemoteDB bool = true
var isDbEnabled bool = true
const sessionKey = "previewToken"

var validToken string

func main() {
	rwPreviewToken := os.Getenv("validToken")
	rwEnv := os.Getenv("RAILWAY_ENVIRONMENT")
	isProduction := rwEnv == "production"
	isLocal := rwEnv == ""

	if rwPreviewToken != "" && !isProduction{
		validToken = rwPreviewToken
		color.Cyan("Preview Token: "+validToken)
	}else {
		validToken = "xyz123"
		color.Cyan("Preview Token: "+validToken)
	}

	getParameters()
	color.Cyan("[ℹ INFO]: Starting *gin* in %s mode", gin.Mode())	

	r := gin.Default()
	
	if !isProduction && !isLocal {
		store := cookie.NewStore([]byte("dsoifjdsla823495jreio89xpgjgftzftttrertertecjipx9f"))
		token := getRandomToken()
		color.Cyan("[ℹ INFO] Session token: "+token)
		r.Use(sessions.Sessions(token, store))
		r.Use(tokenSessionMiddleware())
	}

	
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
		if c.Query("newPage") != "true" && isProduction{
			newsHandler(c)
		}else{
			c.File("public/html/news-testing.html")
		}
	})


	r.POST("/login", func(c *gin.Context){
		POSTloginHandler(c)
	})

	r.POST("/register", func(c *gin.Context) {
		POSTregisterHandler(c)
	})


	r.GET("/connection/info", func(c *gin.Context){
		infoHandler(c)
	})

	r.GET("/versions/:email/:game", func(c *gin.Context){
		if isDbEnabled {
			getVersions(c)
		}else {
			c.String(503, "Service only available with DB enabled!")
		}
	})


	r.POST("/motd", func(c *gin.Context) {
		type MotdRequest struct {
			Message string `json:"message"`
		}

		var req MotdRequest
		if err := c.BindJSON(&req); err != nil {
			c.String(400, "Failed to parse JSON data: %v", err)
			return
		}

		newMotd := req.Message
		currentTime := time.Now()
		success, msg := setMOTD(newMotd, currentTime)
		if !success {
			c.String(500, msg)
			return
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

	r.GET("/ws", wsHandler)
           


	if isLocal{
		color.Magenta("[⚙ RW Environment]: Local development")

	}else{
		color.Magenta("[⚙ RW Environment]: %s", rwEnv)

	}
	color.Green("[✓ SUCCESS] Started Server successfully on http://localhost:8088")
	
	r.Run(":8088")
}