package main

import "fmt"

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	count := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("dari channel 1", data)
			count++
		case data := <-channel2:
			fmt.Println("dari channel 2", data)
			count++
		default:
			fmt.Println("Waiting receive data")
		}
		if count == 2 {
			break
		}
	}
}

func GiveMeResponse(channel chan<- string) {
	channel <- "Hello"
}
