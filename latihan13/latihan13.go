package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("Test MySQL Database")
    db, err := sql.Open("mysql", "admin:passwd@tcp(127.0.0.1:3306)/golang")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

}