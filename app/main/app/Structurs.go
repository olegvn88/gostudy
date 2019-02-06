package main

import "fmt"

type myStructure struct {
	x int
	y int
}

func main(){
	s := myStructure{1, 4}
	p := &s
	p.x = 5
	str := myStructure{y: 2, x: 50}
	fmt.Println(s)
	fmt.Println(str.x)
}