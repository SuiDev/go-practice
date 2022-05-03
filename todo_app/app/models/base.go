package models

import (
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	defer Db.Close()

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		create_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)
}