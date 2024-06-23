package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() *sql.DB {
	Driver := "mysql"
	User := "root"
	Pass := "password123"
	Name := "system"

	connection, err := sql.Open(Driver, User+":"+Pass+"@tcp(127.0.0.1)/"+Name)
	if err != nil {
		log.Fatal(err)
	}
	return connection
}