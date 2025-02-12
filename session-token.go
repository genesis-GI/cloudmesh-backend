package main

import(
	"math/rand"
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func getRandomToken() string {
	b := make([]rune, 16)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func tokenSessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		queryToken := c.Query("previewToken")
		if queryToken != "" {
			session.Set(sessionKey, queryToken)
			session.Save()
		}

		token := session.Get(sessionKey)
		if token == nil || token != validToken {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access Denied"})
			c.Abort()
			return
		}

		c.Set(sessionKey, token)
		c.Next()
	}
}