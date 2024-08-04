package database

import (
	"database/sql"
	"os"
	"the-iron-cygnet/database/sqlc"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Conn    *sql.DB
	Queries *sqlc.Queries
}

var DB = &Database{}

func InitDB() {
	conn, err := sql.Open("sqlite3", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	DB.Conn = conn
	DB.Queries = sqlc.New(DB.Conn)
}
