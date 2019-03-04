package sockets

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {

	listener, _ := net.Listen("tcp", ":5000")

	for {
		conn, err := listener.Accept()
		//conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Can not connect")
			conn.Close()
			continue
		}

		fmt.Println("Connected!")

		bufReader := bufio.NewReader(conn)
		fmt.Println("Start reading")

		go func(c net.Conn) {

			defer conn.Close()

			for {
				rByte, err := bufReader.ReadString('\n')

				if err != nil {
					fmt.Println("Can not read!", err)
					break
				}

				fmt.Printf("- %s", rByte)
				if strings.Trim(rByte, "\r\n") == "exit" {
					conn.Write([]byte("Bye-bye\n"))
					break
				}
			}
		}(conn)

	}
}
