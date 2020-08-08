package dbutils

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

var (
	// DB is an instance of a sqlx DB; in global scope so we can use it throughout the file
	DB  *sqlx.DB
	err error
)

// User is a User table abstraction struct
type User struct {
	UUID        uuid.UUID `db:"uuid" json:"uuid"`
	Username    string    `db:"username" json:"username"`
	Email       string    `db:"email" json:"email"`
	Password    string    `db:"password" json:"password"`
	Description string    `db:"description" json:"description"`
}

// Post is a Post table abstraction struct
type Post struct {
	PostUUID    uuid.UUID `db:"post_uuid" json:"post_uuid"`
	PostCreator uuid.UUID `db:"post_creator" json:"post_creator"`
	PostContent string    `db:"post_content" json:"post_content"`
	Time        time.Time `db:"time" json:"time"`
}

// Open opens the DB
func Open() {
	DB, err = sqlx.Connect("mysql", "root:root@tcp(Awty_Portal:6000)")
	if err != nil {
		panic(err)
	}
}

// SelectUsers gets all users
func SelectUsers(user *User) error {
	return DB.Get(user, "SELECT uuid, username, email, description FROM User")
}

// SelectUserByUUID gets a user by their UUID
func SelectUserByUUID(uuid uuid.UUID, user *User) error {
	return DB.Get(user, "SELECT uuid, username, email, description FROM User where uuid = ?", uuid)
}

// SelectPosts selects all posts
func SelectPosts(post *Post) error {
	return DB.Get(post, "SELECT * FROM Post")
}
