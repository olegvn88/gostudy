package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

type address struct {
	street, city, country string
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.abs())

	fmt.Println(address{"Гарматная", "Киев", "Украина"})
}

func (a address) getAddress() string {
	return a.city + "\n" + a.street + "\n" + a.country

}
