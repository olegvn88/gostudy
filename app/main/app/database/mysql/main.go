package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //driver for data base
)

var (
	db *sql.DB
)

func main() {
	var err error
	// создаем структуру базы данных
	// но соединение просиходит только при первом запросе
	db, err := sql.Open("mysql", "root@tcp(172.19.0.2:3306)/mysql?charset=utf8")
	PanicOnErr(err)

	fmt.Println("Open Connections", db.Stats().OpenConnections)

	// проверяем, что подключение реально произошло (Делаем запроос)
	err = db.Ping()
	PanicOnErr(err)

	fmt.Println("Open Connections", db.Stats().OpenConnections)
	return

	// итерируемся по многим запросам
	//
}

func PanicOnErr(e error) {
	if e != nil {
		panic(e)
	}
}
