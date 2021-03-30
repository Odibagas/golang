package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type mahasiswa struct {
    nim    			string
    nama  			string
    tanggal_lahir	string
    jenis_kelamin	string
}

func connect() (*sql.DB, error) {
    db, err := sql.Open("mysql", "admin:passwd@tcp(127.0.0.1:3306)/golang")
    if err != nil {
        return nil, err
    }

    return db, nil
}

func sqlQueryRow() {
    var db, err = connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()

    var profil = mahasiswa{}
    err = db.
        QueryRow("select COUNT(nim) from mahasiswa").
        Scan(&profil.nim)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Printf("Total NIM: %s\n", profil.nim)
}

func main() {
    sqlQueryRow()
}