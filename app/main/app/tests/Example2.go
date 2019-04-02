package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r circle) area() float64 {
	return r.radius * math.Pi
}

func (r rectangle) perim() float64 {
	return 2 * (r.width + r.height)
}

func (r circle) perim() float64 {
	return 2 * math.Pi * r.radius
}

func calculate(g geometry) {
	fmt.Println(g)
	fmt.Println("Area is", g.area())
	fmt.Println("Perimetr is", g.perim())
}

func main() {
	rect := rectangle{width: 4, height: 5}
	circ := circle{radius: 3}
	calculate(rect)
	calculate(circ)
}
