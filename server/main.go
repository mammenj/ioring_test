package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go serve(conn)
	}

}

func serve(connection net.Conn) {
	defer connection.Close()
	for {
		buffer := make([]byte, 1024)
		_, err := connection.Read(buffer[:])
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(buffer))
	}
}
