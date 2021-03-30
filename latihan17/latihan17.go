package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type mahasiswa struct {
	nim           string
	nama          string
	tanggal_lahir string
	jenis_kelamin string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "admin:passwd@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("update mahasiswa set nama = ? where nim = ?", "Gojo", "3430301322")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success!")
}

func main() {
	sqlExec()
}
