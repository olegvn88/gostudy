package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //driver for data base
)

var (
	database *sql.DB
)

type RowsList struct {
	fio, info string
	id, score int
}

func main() {
	var err error
	database, err := sql.Open("mysql", "root:root@tcp(172.19.0.2)/testbase")
	PanicForError(err)

	rows, err := database.Query("select fio, score from students")

	for rows.Next() {
		var fio string
		var score string
		err = rows.Scan(&fio, &score)
		PanicForError(err)
		fmt.Println(fio, score)
	}
	rows.Close()

	var r RowsList
	err = database.QueryRow("select fio from students").Scan(&r.fio)
	PanicForError(err)

	fmt.Printf("fio = %s", r.fio)
}

func PanicForError(e error) {
	if e != nil {
		panic(e)
	}
}
