package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() []int {
	//line := 0, 1, 1, 2, 3, 5, 8, 13, 21
	var k int
	fib := func() []int {
		arr := make([]int, 10)
		for k = 2; k < len(arr); k++ {
			arr[0] = 0
			arr[1] = arr[0] + 1
			arr[k] = arr [k - 1] + arr[k - 2]
			//fmt.Print(arr[k], ", ")
		}
		return arr
	}
	return fib
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
	fmt.Println(f())
	}
}
