package main

import (
	"fmt"
	"reflect"
)

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

	updateSlice()
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

func makeSLice() {
	a := make([]int, 15)
	printMakeSlice("a", a)

	b := make([]int, 0, 15)
	printMakeSlice("b", b)

	c := b[:2]
	printMakeSlice("c", c)

	d := c[2:5]
	printMakeSlice("d", d)
}

func printMakeSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func updateSlice() {
	var s1 []int
	printNewSlice(s1)
	s1 = append(s1, 12)
	s1 = append(s1, 12)
	s1 = append(s1, 12)
	s1 = append(s1, 12)
	s1 = append(s1, 12)
	var s2 []int
	s2 = append(s2, 3, 4)
	s1 = append(s1, 40)
	s2 = append(s2, 12)
	s1 = append(s1, s2...)
	printNewSlice(s1)

	//Создать слайс нужной длины
	slice3 := make([]int, 10, 20)
	slice3 = append(slice3, 2, 20, 12)
	slice3[1] = 40
	slice3[2] = 30
	slice4 := make([]int, len(slice3), cap(slice3))
	copy(slice4, slice3)
	slice4[0] = 9999
	printNewSlice(slice3)
	printNewSlice(slice4)

	arr := []int{2, 4, 56}
	slice5 := arr[:]
	slice5 = append(slice5, 3, 5, 5, 10, 40)
	arr[0] = 100
	printNewSlice(slice5)
	fmt.Printf("slice elemenst are  %v\nnewline\n", slice5)
	//printNewSlice(slice4, slice5)
}

func printNewSlice(testSlice []int) {
	fmt.Println("======", reflect.TypeOf(testSlice), "======")
	fmt.Println("slice contains ", testSlice)
	fmt.Println("length of slice", len(testSlice))
	fmt.Println("capacity of slice", cap(testSlice))
	fmt.Println("-------------------------")
}
