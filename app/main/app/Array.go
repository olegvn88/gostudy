package main

import "fmt"

var arr [3]string

func main() {
	arr[0] = "Say"
	arr[1] = "Hi"
	fmt.Println(arr)
	fmt.Println(arr[0])

	arr1 := [4]string{"Hi", "Dear", "Friend"}
	fmt.Println(arr1[0])
	fmt.Println(arr1[1])
	fmt.Println(arr1[2])
	fmt.Println(arr1[3])
	fmt.Println(cap(arr))
}
