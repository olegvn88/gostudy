package main

import "fmt"

type Vertex struct {
	lat, long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	//fmt.Println(m["Bell Labs"])
	putStructureToMap()
}

var m2 = map[int]persons{
	//"Bell Labs": {40, 60},
	//"Google":    {36, -122},
	1: {"Oleg", 60},
	2: {"Nastya", -122},
}

type persons struct {
	name string
	age  int
}

func putStructureToMap() {
	fmt.Println(m2 ,"\n")
	//fmt.Println(Vertex{18,39})
	//fmt.Println(persons{"Oleg", 30})

	s := persons{"Dima", 32}

	s = persons{"Nastya", 27}
	fmt.Println(s)

	sp := &s
	sp.name = "Alex"
	sp.age = 40

	fmt.Println(s.age)
	fmt.Println(sp.age)
	fmt.Println(s)

	str := []struct {
		a string
		b int
	}{
		{"Oleg", 3}}

	fmt.Println(str)
}
