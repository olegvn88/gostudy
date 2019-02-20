package main

import (
	"fmt"
	"os"
)

//print 9 8 7 6 5 4 3 2 1 0
func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Print(i, " ")
	}
}

//defer example of using
func ReadFile(f string) ([]byte, error) {
	file, err := os.OpenFile(f, 0, 0666)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	return nil, nil
}
