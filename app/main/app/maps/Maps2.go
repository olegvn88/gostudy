package main

import (
	"fmt"
)

func main() {
	var map1 map[string]string
	printMyMap(map1)

	map2 := map[string]string{}
	map2["test"] = "ok"
	fmt.Println(map2)

	var map3 = make(map[string]string)
	map3["age"] = "30"
	printMyMap(map3)

	someValue := map3["test"]
	fmt.Println("Random value", someValue, len(map3))

	lastName, ok := map3["test"]
	fmt.Println("Key is ", lastName, "exist:", ok)

	_, ok = map3["test"]
	fmt.Println("Key exists:", ok)

	_, ok = map2["test"]
	fmt.Println("Key exists:", ok)

	delete(map2, "test")
	_, exist := map2["test"]

	fmt.Println("Key exists", exist)

	key := "First Name"

	map4 := map2
	map4[key] = "test"
	fmt.Println(map4, map2)

	map5 := map[string]string{
		"firstName": "Oleg",
		"lastName":  "Nesterov",
	}

	//Проверка на существование ключа
	if _, ok := map5["firstName"]; ok {
		println("firstName key exist = ", ok)
	} else {
		println("no firstName")
	}

	if firstName, ok := map5["firstName"]; !ok {
		println("no firstName")
	} else if firstName == "Oleg" {
		println("First name is Oleg")
	} else {
		println("first name is not Oleg")
	}

}

func printMyMap(inputMap map[string]string) {
	fmt.Println("->", inputMap)
}
