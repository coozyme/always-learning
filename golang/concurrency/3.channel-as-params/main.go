package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)

	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel //reciew

	fmt.Println("d", data)

	time.Sleep(2 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "from func GiveMeResponse" //insert data to channel
}
