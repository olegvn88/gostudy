package main

import "fmt"

func sum(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}

func main() {
	result := []int{4, 56, 2, 2, 5}
	fmt.Println(result, sum(result...)) // ... распаковывает слайс в одиночные аргументы, которые передаются функции

	//anonimus function
	func(line string) {
		fmt.Println(line)
	}("Hello")

	//you can assign function to the variable
	printLine := func(line string) {
		fmt.Println(line)
	}
	printLine("Oleg")

	type strFuncType func(string) string

	worker := func(f strFuncType) {
		f("Hi")
	}
	fmt.Println(worker)
}
