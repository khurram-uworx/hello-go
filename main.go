package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // will listen on 0.0.0.0:8080
}
