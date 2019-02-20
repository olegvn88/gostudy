package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	tempValue := bufio.NewScanner(os.Stdin)
	var result []int
	var i = 0
	for {
		tempValue.Scan()
		if tempValue.Text() != "exit" {
			val, _ := strconv.Atoi(tempValue.Text())
			result = append(result, val)
			i++
		} else {
			break
		}
	}
	fmt.Println(printSum(result...))
}

func printSum(stuff ...int) (res int) {
	for i, _ := range stuff {
		res += stuff[i]
	}
	return
}

func init() {
	fmt.Println("Hello")
}
