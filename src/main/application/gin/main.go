package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// go get -u github.com/gin-gonic/gin
func main() {
	fmt.Println("run server begin :: ")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
