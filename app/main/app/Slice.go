package main

import "fmt"

func main() {
	//s := []int{2, 3, 5, 7, 11, 13}
	//
	//s = s[1:4]
	//fmt.Println(s)
	//
	//s = s[:2]
	//fmt.Println(s)
	//
	//s = s[1:]
	//fmt.Println(s)


	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	str := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(str)
	fmt.Println()
	fmt.Println()
	slice2()
}

func slice2() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	fmt.Println("===============")
	fmt.Println("===============")
	fmt.Println("===============")
	makeSLice()
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func makeSLice(){
	a := make([]int, 15)
	printMakeSlice("a", a)

	b := make([]int ,0, 15)
	printMakeSlice("b", b)

	c := b[:2]
	printMakeSlice("c", c)

	d := c[2:5]
	printMakeSlice("d", d)
}

func printMakeSlice(s string, x []int)  {
	fmt.Printf("%s len=%d cap=%d %v\n", s , len(x), cap(x), x)
}