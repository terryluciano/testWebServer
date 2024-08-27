package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println(ctx.ClientIP())
		ctx.Next()
	}
}

func main() {
	r := gin.New()

	r.Use(Logger())

	r.GET("ping", func(c *gin.Context) {
		fmt.Println(c.ClientIP())
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	r.Run(":4000")
}
