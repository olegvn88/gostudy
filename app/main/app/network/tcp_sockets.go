package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	listener, _ := net.Listen("tcp", ":5000")

	for {
		// Ждем пока не прийдет наш клиент
		conn, err := listener.Accept()

		if err != nil {
			fmt.Printf("Can not connect!")
			conn.Close()
			continue
		}
		fmt.Println("Connected!")

		//создаем ридер для чтения информации из сокета

		bufReader := bufio.NewReader(conn)
		fmt.Println("Start reading")

		go func(c net.Conn) {

			defer conn.Close()

			for {
				//побайтово читаем
				rByte, err := bufReader.ReadByte()

				if err != nil {
					fmt.Println("Can not read!", err)
					break
				}
				//выводим то, что написали в консоле
				fmt.Print(string(rByte))
				conn.Write([]byte("Received!"))
			}
		}(conn)
	}
}
