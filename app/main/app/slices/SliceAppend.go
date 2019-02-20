package slices

import (
	"fmt"
)

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	fmt.Println("============== \n")

	a := make([]int, 10)
	for i := 0; i < 5; i++ {
		a[i] = i
	}
	printSliceAppends("a", a)
	arr2 := append(a, 2000, 1000, 3000, 5, 6, 7, 44, 20,4,2,56,4,4)
	fmt.Println("append = ", arr2)
	fmt.Println("append = ", cap(arr2))
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceAppends(s string, arr []int) {
	fmt.Printf("%s len = %d cap = %d %d \n", s, len(arr), cap(arr), arr)
}
