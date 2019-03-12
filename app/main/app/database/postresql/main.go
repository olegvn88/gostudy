package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //driver for data base
	_ "github.com/lib/pq"
	"github.com/olegvn88/gostudy/app/main/app/database/properties"
)

var (
	db   *sql.DB
	prop properties.PropList
)

// PrintByID print student by id
func PrintByID(id int64) {
	var fio string
	var info sql.NullString
	// var info string
	var score int
	row := db.QueryRow("SELECT fio, info, score FROM students WHERE id = $1", id)
	// fmt.Println(row)
	err := row.Scan(&fio, &info, &score)
	PanicForErrors(err)
	fmt.Println("PrintByID:", id, "fio:", fio, "info:", info, "score:", score)
}

func main() {
	var err error
	var openLink string
	openLink = "user=" + prop.GetPostgressLogin() + " password=" + prop.GetPostgressPassword() + " host=" + prop.GetPostgresAddress() + " dbname=" + prop.GetDatabaseName() + " sslmode=disable"
	fmt.Println(openLink)
	// создаем структуру базы данных
	// но соединение просиходит только при первом запросе
	//db, err = sql.Open("postgres", "user=postgres password=root host=172.19.0.3 dbname=testbase sslmode=disable")
	db, err = sql.Open("postgres", openLink)
	PanicForErrors(err)

	//db.SetMaxOpenConns(10)

	//fmt.Println("Open Connections", db.Stats().OpenConnections)

	// проверяем, что подключение реально произошло (Делаем запроос)
	err = db.Ping()
	PanicForErrors(err)

	//fmt.Println("Open Connections", db.Stats().OpenConnections)

	// итерируемся по многим записям
	// Exec исполняет запрос и возвращает записи

	rows, err := db.Query("SELECT fio from students")
	PanicForErrors(err)

	for rows.Next() {
		var fio string
		err = rows.Scan(&fio)
		PanicForErrors(err)

		fmt.Println("rows.Next fio: ", fio)
	}
	// надо закрыть соединение, иначе будет течь, а их всего 10
	rows.Close()
	//return

	var fio string
	row := db.QueryRow("SELECT fio from students where id = 1")
	err = row.Scan(&fio)
	PanicForErrors(err)
	fmt.Println("db.QueryRow fio: ", fio)
	//return

	// Exec испролняет запрос и возвращает сколько строк было затронуто и последний ID вставленной записи
	// Символ $1 является placeholder-ом, все последующие значения атов-экранируются и подставляются с правильной

	var lastID int64
	err = db.QueryRow(
		"INSERT INTO students (fio, score) VALUES ($1, 0) RETURNING id",
		"Ivan Ivanov",
	).Scan(&lastID)
	PanicForErrors(err)

	fmt.Println("Insert - lastInsertId: ", lastID)

	PrintByID(lastID)
	return
}

func PanicForErrors(e error) {
	if e != nil {
		panic(e)
	}
}
