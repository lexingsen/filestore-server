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
	dbx, err := sql.Open("mysql", "root:111111@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
	if err != nil {
		fmt.Println("Failed to connet to the mysql, err:" + err.Error())
		os.Exit(1)
	}
	log.Printf("mysql is opened!")
	dbx.SetMaxOpenConns(1000)
	err = dbx.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
	log.Printf("mysql ping is pong!")
	db = dbx
}

func DBConn() *sql.DB {
	return db
}
