package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	listener, _ := net.Listen("tcp", ":5000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Cannot connect")
			conn.Close()
			continue
		}

		fmt.Println("Connected")

		buferReader := bufio.NewReader(conn)
		fmt.Println("Start reading")

		go func(c net.Conn) {
			defer conn.Close()

			for {

				rByte, err := buferReader.ReadString('\n')

				if err != nil {
					fmt.Println("Cannot read!")
					break
				}

				fmt.Println(rByte)
				conn.Write([]byte("Bye"))

			}
		}(conn)

	}

}
