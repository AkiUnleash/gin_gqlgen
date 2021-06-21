package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(sampleMiddleware())
	r.GET("/", func(c *gin.Context) {
		log.Println("main logic")
		c.JSON(200, gin.H{"message": "Hello!"})
	})
	r.Run()
}

func sampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before logic")
		c.Next()
		log.Println("after logic")
	}
}
