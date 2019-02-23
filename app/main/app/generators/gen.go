package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//c := greeting("Hey")

	c := funIn(greeting("Hey!"), byeChan("Bye!"))
	for i := 0; i < 10; i++ {
		fmt.Printf("She says: %q\n", <-c)
	}
}

func greeting(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			time.After(30)
		}
	}()
	return c
}

func byeChan(msg string) <-chan string {

	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		}
	}()
	return c
}

func funIn(input1, input2 <-chan string) <-chan string {

	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
