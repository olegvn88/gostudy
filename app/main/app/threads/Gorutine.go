package main

import (
	"fmt"
)

func main() {
	closure()
	return
	fmt.Printf("start")
	//можно запустить функцию
	go process(0)
	//можно запустить анонимную йункцию
	go func() {
		fmt.Println("Анонимный запуск")
	}()

	//можем запустить много горутин
	for i := 0; i < 1000; i++ {
		go process(i)
	}
	// нужно дождаться завершения выполнения

}

func process(i int) {
	fmt.Println("Обработка: ", i)
}

func closure() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("Got", i)
		}(i) // анонимная функция
	}
	fmt.Scanln()
}
