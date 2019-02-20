package main

import (
	"fmt"
)

type myFloat float64

func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := myFloat(2503)
	fmt.Println(v.abs())
}
