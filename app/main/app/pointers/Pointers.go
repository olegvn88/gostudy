package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i                                     // point to i
	fmt.Println(*p)                             // read i through the pointer
	*p = 21                                     // set i through the pointer
	fmt.Println("// see the new value of i", i) // see the new value of i

	p = &j // point to j
	fmt.Println(*p)
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	fmt.Println("------")
	printPointer()
}

func printPointer() {
	i := 5
	b := &i
	//var b *int
	*b = 6
	fmt.Println(i)

	sl := []int{0, 0, 0}
	fmt.Println(sl)
	sum(sl)
	fmt.Println(sl)
}

func sum(items []int) {
	items[2] = 1
}
