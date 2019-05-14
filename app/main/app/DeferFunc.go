package main

import "fmt"

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend:", err)
		}
	}()
	fmt.Println("Some userful work")
	panic("something bad happen")
	return
}

func main() {
	deferTest()
	return
}
