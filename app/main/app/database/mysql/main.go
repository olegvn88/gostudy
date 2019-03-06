package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //driver for data base
)

var (
	db *sql.DB
)

// PrintByID print student by id
func PrintByID(id int64) {
	var fio string
	var info sql.NullString //
	// var info string // сюда запишется значение по умолчанию
	var score int
	row := db.QueryRow("SELECT fio, info, score FROM students WHERE id = ?", id)
	// fmt.Println(row)
	err := row.Scan(&fio, &info, &score)
	PanicOnErr(err)
	fmt.Println("PrintByID:", id, "fio:", fio, "info:", info, "score:", score)
}

func main() {
	var err error
	// создаем структуру базы данных
	// но соединение просиходит только при первом запросе
	db, err := sql.Open("mysql", "root:root@tcp(172.19.0.2:3306)/olnester?charset=utf8&interpolateParams=true")
	PanicOnErr(err)

	db.SetMaxOpenConns(10)

	fmt.Println("Open Connections", db.Stats().OpenConnections)

	// проверяем, что подключение реально произошло (Делаем запроос)
	err = db.Ping()
	PanicOnErr(err)

	fmt.Println("Open Connections", db.Stats().OpenConnections)

	// итерируемся по многим записям
	// Exec исполняет запрос и возвращает записи

	rows, err := db.Query("SELECT fio, score from students")
	PanicOnErr(err)

	for rows.Next() {
		var fio string
		var score int
		err = rows.Scan(&fio, &score)
		PanicOnErr(err)

		fmt.Println("rows.Next fio: ", fio, "score: ", score)
	}
	// надо закрыть соединение, иначе будет течь, а их всего 10
	rows.Close()
	//return

	var fio string
	row := db.QueryRow("SELECT fio from students where id = 1")
	err = row.Scan(&fio)
	PanicOnErr(err)
	fmt.Println("db.QueryRow fio: ", fio)
	//return

	// Exec испролняет запрос и возвращает сколько строк было затронуто и последний ID вставленной записи
	// Символ ? является placeholder-ом, все последующие значения атов-экранируются и подставляются с правильной

	result, err := db.Exec(
		"INSERT INTO students (`fio`,`score`) values (?, 0)", // запрос
		"Oleg Vladislavovich",                                // значение
	)
	//запрос "INSERT INTO students (`fio`,`score`) values (?, 0)" распарсится и будет ждать уже следующего запроса с данными "Oleg Vladislavovich"
	PanicOnErr(err)

	affected, err := result.RowsAffected()
	PanicOnErr(err)
	lastID, err := result.LastInsertId()
	PanicOnErr(err)
	fmt.Println("Insert = RowsAffected", affected, "LastInsertID: ", lastID)

	//PrintByID(lastID)

	result, err = db.Exec(
		"UPDATE students SET info = ?  where id = ?",
		"test user",
		lastID,
	)
	PanicOnErr(err)

	affected, err = result.RowsAffected()
	PanicOnErr(err)
	fmt.Println("Update - RowsAffected", affected)

	//PrintByID(lastID)
	//return

	// использование prepared statements
	// Prepare подготавливает запись (шлет запрос на сервере, там он парсится и готов принимать данные
	stmt, err := db.Prepare("UPDATE students SET info = ?, score = ? where id = ?")
	PanicOnErr(err)
	// Exec ля prepares statement отправляет данные на сервер - там запрос уже распаршен, только исполнение
	result, err = stmt.Exec("prepared statement update", lastID, lastID)
	PanicOnErr(err)
	result, err = stmt.Exec("8 prepared statement update", lastID, 8)
	PanicOnErr(err)

	//return

	affected, err = result.RowsAffected()
	PanicOnErr(err)
	fmt.Println("Update - RpwsAffected", affected)

}

func PanicOnErr(e error) {
	if e != nil {
		panic(e)
	}
}
