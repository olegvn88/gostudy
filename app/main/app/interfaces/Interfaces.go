package main

import (
	"fmt"
	"math"
)

type abser interface {
	abs() float64
}

func main() {
	var a abser
	f := myFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	fmt.Println(a.abs())
}

type myFloat float64

func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	x, y float64
}

func (v *Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
