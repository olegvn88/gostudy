package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //driver for data base
	"github.com/jinzhu/gorm"
)

type TestUser struct {
	ID       int
	Username string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/testbase")

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	db.CreateTable(&TestUser{})

	fmt.Println("table created", db.HasTable(&TestUser{})) //true

	db.DropTable(&TestUser{})

	defer db.Close()

}
