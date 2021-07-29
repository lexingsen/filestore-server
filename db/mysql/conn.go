package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:111111@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
	if err != nil {
		fmt.Println("Failed to connet to the mysql, err:" + err.Error())
		os.Exit(1)
	}
	log.Printf("mysql is opened!")
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
	log.Printf("mysql ping is pong!")
}

func DBConn() *sql.DB {
	return db
}
