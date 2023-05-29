package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

const (
	port            = 8080
	numberOfClients = 10000
)

var testData = []byte("Hello world")

func main() {
	//time.Sleep(time.Second * 2)
	now := time.Now()
	wg := &sync.WaitGroup{}
	//c := make(chan string) // We don't need any data to be passed, so use an empty struct

	for i := 0; i < numberOfClients; i++ {

		wg.Add(1)

		go func(j int) {
			//time.Sleep(time.Millisecond * 10)
			defer wg.Done()
			// defer func() {
			// 	c <- "Client: " + strconv.Itoa(j) // sgnal that the routine has completed

			// }()
			//fmt.Println("Client: ", i)

			//conn, err := net.DialTimeout("tcp", fmt.Sprintf("127.0.0.1:%d", port), time.Second)
			conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))

			if err != nil {
				log.Panic(err)
			}

			_, err = conn.Write(testData)
			if err != nil {
				log.Panic()
			}

			// if n != len(testData) {
			// 	log.Panic()
			// }

			//buffer := make([]byte, len(testData))

			//_, err = conn.Read(buffer)
			//if err != nil {
			//	log.Panic()
			//}
			//defer conn.Close()

			// if n != len(testData) {
			// 	log.Panic()
			// }

		}(i)
	}
	wg.Wait()
	// for i := 0; i < numberOfClients; i++ {
	// 	fmt.Println("Completed", <-c)
	// }

	fmt.Println("Time taken:", time.Since(now), "for", numberOfClients, "clients")
}
