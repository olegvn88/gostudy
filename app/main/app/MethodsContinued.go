package main

import (
	"fmt"
	"math"
)

type myFloat float64

func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := myFloat(-math.Sqrt2)
	fmt.Println(v.abs())
}
