package main

import "fmt"

var c chan int

func main() {
	//create chanel
	c := make(chan string)

	//start writting go rutine
	go great(c)

	for i := 0; i < 5; i++ {
		// read a few line from the chanel
		fmt.Println(<-c, ",", <-c)
	}

	//stuff := make(chan int, 7)
	//for i := 0;i < 19 ; i++  {
	//	stuff <- i
	//}
	//close(stuff)
	//fmt.Println("Res", process(stuff))
}

func process(input <-chan int) (res int) {
	for r := range input {
		res += r
	}
	return
}

func great(c chan<- string) {
	// run unlimited loop
	for {
		// and write a few lines to the chanel
		//Subprogram will be disabled until somebody want to read it
		c <- fmt.Sprintf("Владыка")
		c <- fmt.Sprintf("Штурмовик")

	}
}
