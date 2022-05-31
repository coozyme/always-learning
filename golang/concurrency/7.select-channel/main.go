package main

import "fmt"

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// for {
	select {
	case data := <-channel1:
		fmt.Println("dari channel 1", data)
	case data := <-channel2:
		fmt.Println("dari channel 2", data)
	}
	// }
}

func GiveMeResponse(channel chan<- string) {
	channel <- "Hello"
}
