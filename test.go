// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	db, err := sql.Open("mysql", "root:111111@tcp(127.0.0.1:3306)/test?charset=utf8")
// 	if err != nil {
// 		fmt.Println("Failed to connet to the mysql, err:" + err.Error())
// 		os.Exit(1)
// 	}
// 	log.Printf("mysql is opened!")
// 	db.SetMaxOpenConns(1000)
// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Println("Failed to connect to mysql, err:" + err.Error())
// 		os.Exit(1)
// 	}
// 	log.Printf("mysql ping is pong!")

// 	stmt, err := db.Prepare("insert into user(id,name,age,sex) values (?,?,?,?)")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("==============================================================================")
// 	if err != nil {
// 		fmt.Println("Failed to prepare statement, err:" + err.Error())
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(10, "lsc", 30, 2)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }
