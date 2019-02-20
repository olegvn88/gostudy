package main

import "fmt"

type example struct {
	Flag    bool
	counter int
	pi      float64
}

func main() {

	var ex3 example
	fmt.Println(ex3)
	ex3 = example{
		Flag:    true,
		counter: 10,
		pi:      3.141592,
	}
	fmt.Println(ex3)

	ex3 = example{false, 30, 4.4}
	fmt.Println(ex3)

	ex3 = example{}
	fmt.Println(ex3)

}
