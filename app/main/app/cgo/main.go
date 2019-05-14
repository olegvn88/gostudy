package main

import "time"

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(60 * 4) // 4 minutes
		}()
	}
	time.Sleep(5 * time.Minute)
}
