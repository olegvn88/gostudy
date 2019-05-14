package main

import "fmt"

/*
	The range form of the for loop iterates over a slice or map.
	When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
*/

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	str := "Nesterov"
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	fmt.Println("--------------------\n")
	for pos, char := range str {
		fmt.Printf("position: %d char: %c\n", pos, char)
	}
	fmt.Println("--------------------\n")
	printRange()

}

func printRange() {
	pow := make([]int, 4)
	for i := range pow {
		fmt.Println("i - ", i)
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
