package main

import (
	"fmt"
	"strconv"
)

func compute(fn1 func(float64, float64) int) int {
	return fn1(3, 4)
}

func main() {
	//hypot := func(x, y float64) float64 {
	//	return math.Sqrt(x*x + y*y)
	//}
	//fmt.Println(hypot(5, 12))
	//
	//fmt.Println(compute(hypot))
	//fmt.Println(compute(math.Pow))

	fullName := func(n string, a int) string {
		return "My name is " + n + " I am " + strconv.Itoa(a) + " is old"
	}

	fmt.Println(fullName("Dima", 4))
	fmt.Println(printNameAndAge(fullName))

	myFunc := func(g, h float64) int {
		s := "3"
		a, _ := strconv.Atoi(s)

		return a
	}

	fmt.Println(compute(myFunc))
}

func printNameAndAge(humanFunction func(string, int) string) string {
	return humanFunction("Oleg", 30)
}
