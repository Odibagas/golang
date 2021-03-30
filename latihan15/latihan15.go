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

func sqlQuery() {
    db, err := connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()

    rows, err := db.Query("select nim, nama, tanggal_lahir, jenis_kelamin from mahasiswa")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer rows.Close()

    var profil []mahasiswa

    for rows.Next() {
        var each = mahasiswa{}
        var err = rows.Scan(&each.nim, &each.nama, &each.tanggal_lahir, &each.jenis_kelamin)

        if err != nil {
            fmt.Println(err.Error())
            return
        }

        profil = append(profil, each)
    }

    if err = rows.Err(); err != nil {
        fmt.Println(err.Error())
        return
    }

    for _, each := range profil {
        fmt.Println(each.nim, each.nama, each.tanggal_lahir, each.jenis_kelamin)
    }
}

func main() {
    sqlQuery()
}