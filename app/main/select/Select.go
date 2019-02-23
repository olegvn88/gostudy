package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- "from 1"
		time.Sleep(time.Second * 2)
	}()

	go func() {
		c2 <- "from 2"
		time.Sleep(time.Second * 3)
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
				fmt.Println(time.Now().Format(time.ANSIC))
				fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout", time.Now().Local().Weekday())
			}
		}
	}()
	fmt.Scanln()
}
