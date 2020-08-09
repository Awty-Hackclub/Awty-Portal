package handler

import (
	"Awty-Portal/backend/dbutils"
	"github.com/gin-gonic/gin"
)

// Users returns a json result of the all users in the User table
func Users(c *gin.Context) {
	users := []dbutils.User{}
	dbutils.SelectUsers(users)
}

// User returns a json result of a User's data (with desired UUID)
func User(c *gin.Context) {

}

// Posts returns a json result of all the posts in the Post table
func Posts(c *gin.Context) {

}

// Post returns a json result of a post by post_uuid
func Post(c *gin.Context) {
	
}
