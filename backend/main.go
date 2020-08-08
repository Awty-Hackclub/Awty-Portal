package main

import (
	"Awty-Portal/backend/handler"
	"Awty-Portal/backend/dbutils"
	"github.com/gin-gonic/gin"
)

func init() {
	dbutils.Open()
}

func main() {
	r := gin.Default()
	r.GET("/users", handler.Users)

	r.GET("/user", handler.User)

	r.GET("/posts", handler.Posts)

	r.GET("/post", handler.Post)
	r.Run()
}
