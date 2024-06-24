package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() *sql.DB {
	Driver := "mysql"
	User := os.Getenv("DB_USER")
	Pass := os.Getenv("DB_PASSWORD")
	Name := os.Getenv("DB_NAME")
	Host := os.Getenv("DB_HOST")
	Port := os.Getenv("DB_PORT")

	connection, err := sql.Open(Driver, User+":"+Pass+"@tcp("+Host+":"+Port+")/"+Name)
	if err != nil {
		log.Fatal(err)
	}
	return connection
}

// func ConnectionDB() *sql.DB {
// 	Driver := "mysql"
// 	User := "root"
// 	Pass := "password123"
// 	Name := "system"

// 	connection, err := sql.Open(Driver, User+":"+Pass+"@tcp(127.0.0.1)/"+Name)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return connection
