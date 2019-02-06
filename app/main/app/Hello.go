package main

import (
	"fmt"
)

var name string

func add(x, y int, name string) (int, string, string) {
	var welcome string
	name = name
	if (name == "Oleg") || (name == "Dima") {
		welcome = name
		return x + y, welcome, name
	} else {
		return 0, "Unknown", "Unknown"
	}
}

func main() {
	fmt.Println(add(42, 13, "Dima"))
	fmt.Println(add(30, 13, "Oleg"))
	fmt.Println(swap(1, 5))
	x, y := swap(1, 5)
	fmt.Println("x = ", x, "y = ", y)
	fmt.Printf("Split result = %T, size of", split(5, 2))

}

func swap(number1, number2 int) (int, int) {
	return number2, number1
}

func split(x, y float32) float32 {
	z := x / y
	return z
}