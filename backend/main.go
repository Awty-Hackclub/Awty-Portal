package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {

	})
	r.GET("/user", func(c *gin.Context) {

	})
	r.GET("/posts", func(c *gin.Context) {

	})
	r.GET("/post", func(c *gin.Context) {

	})
}
