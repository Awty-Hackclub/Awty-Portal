package dbutils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	// DB is an instance of a sqlx DB; in global scope so we can use it throughout the file
	DB  *sqlx.DB
	err error
)

// Open opens the DB
func Open() {
	DB, err = sqlx.Connect("mysql", "root:root@tcp(Awty_Portal:6000)")
	if err != nil {
		panic(err)
	}
}
