package main

import(
	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context){
	c.String(200, "comming soon!")
}