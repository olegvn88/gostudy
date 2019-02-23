package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	die := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	res1 := boring1(&wg, die)
	res2 := boring1(&wg, die)

	for i := 0; i < 5; i++ {
		fmt.Printf("1st say: %q\n", <-res1)
		fmt.Printf("1st say: %q\n", <-res2)
	}

	die <- true
	//Ждем пока все горутины закончат выполняться
	wg.Wait()
}

func boring1(wg *sync.WaitGroup, die chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- fmt.Sprintf("boring %d", rand.Intn(100)):
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			case <-die:
				fmt.Println("Jobs done!")
				wg.Done()
				return
			}
		}
	}()
	return c
}
