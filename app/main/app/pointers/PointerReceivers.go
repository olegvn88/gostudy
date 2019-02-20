package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

func (v vertex) abs() float64 {
	fmt.Println("v.x = ", v.x)
	fmt.Println("v.y = ", v.y)
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *vertex) scale(f float64) {
	v.x = v.x * f
	//fmt.Println(v.x)
	//v.y = v.y * f
}

func main() {
	v := vertex{3, 4}
	v.scale(10)
	fmt.Println(v.abs())
}
