package main

import (
	"fmt"
	"math"
)

type myStruct struct {
	x, y float64
}

func abs(s myStruct) float64 {
	return math.Sqrt(s.x*s.x + s.y*s.y)
}

func main() {
	v := myStruct{3, 4}
	fmt.Println(abs(v))
}
