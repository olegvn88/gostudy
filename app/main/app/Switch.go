package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0;i < 10 ; i++ {
		defer fmt.Println("Hello world")
		getCurrentTime()
		time.Sleep(time.Second)
	}
}

func getKindOfOS() string {
	result := runtime.GOOS
	return result
}

func getCurrentTime(){
	fmt.Println(time.Now().Hour(),":", time.Now().Minute(), ":", time.Now().Second())
}
