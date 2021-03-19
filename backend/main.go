package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/url-shortener/backend/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("loading .env file failed: %v\n", err)
	}

	r := gin.Default()
	r.Use(corsMiddleware())

	r.POST("/encode", handlers.Encode)
	r.GET("/:id", handlers.Redirect)
	r.GET("/:id/info", handlers.Info)

	r.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))

}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
