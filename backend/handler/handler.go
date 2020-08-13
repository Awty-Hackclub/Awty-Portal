package handler

import (
	"Awty-Portal/backend/dbutils"
	"errors"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Users returns a json result of the all users in the User table
func Users(c *gin.Context) {
	users := []dbutils.User{}
	dbutils.SelectUsers(users)
	c.JSON(200, users)
}

// User returns a json result of a User's data (with desired UUID)
func User(c *gin.Context) {
	uuidArg := c.Query("uuid")
	if uuidArg == "" {
		c.AbortWithError(400, errors.New("You need to request with a uuid query string"))
	}
	user := dbutils.User{}
	conversion, err := uuid.FromString(uuidArg)
	if err != nil {
		c.AbortWithError(400, errors.New("Invalid UUID"))
	}
	dbutils.SelectUserByUUID(conversion, &user)
	c.JSON(200, user)
}

// Posts returns a json result of all the posts in the Post table
func Posts(c *gin.Context) {
	posts := []dbutils.Post{}
	dbutils.SelectPosts(posts)
	c.JSON(200, posts)
}

// Post returns a json result of a post by post_uuid
func Post(c *gin.Context) {
	uuidArg := c.Query("post_uuid")
	if uuidArg == "" {
		c.AbortWithError(400, errors.New("You need to request with a uuid query string"))
	}
	post := dbutils.Post{}
	conversion, err := uuid.FromString(uuidArg)
	if err != nil {
		c.AbortWithError(400, errors.New("Invalid UUID"))	
	}
	dbutils.SelectPostsByUUID(conversion, &post)
	c.JSON(200, post)
}
