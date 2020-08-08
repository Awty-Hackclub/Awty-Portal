package dbutils

import (
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

// Open opens the DB
func Open() {
	DB, err = sqlx.Connect("mysql", "root:root@tcp(Awty_Portal:6000)")
	if err != nil {
		panic(err)
	}
}
