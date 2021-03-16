package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/url-shortener/backend/handlers"
)

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())

	r.POST("/encode", handlers.Encode)
	r.GET("/:url", handlers.Redirect)
	r.GET("/:url/info", handlers.Info)

	r.Run()

}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
