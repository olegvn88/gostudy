package main

import "fmt"

func main() {
	a, b := 1, 2
	result := level1(a, b)
	fmt.Println(result)
}

func level1(c int, d int) int {
	e := 3
	return level2(c, d, e)
}

func level2(f int, g int, h int) int {
	return f + g + h
}
