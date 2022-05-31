package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}

/*
	Cara menerima dan mengirim data pada channel
	Menerima / Memasukan data =  channel <- "Channel masuk"
	Mengirim data =  <-channel
*/

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Channel masuk"
}

func OnlyOut(channel <-chan string) {
	time.Sleep(2 * time.Second)
	data := <-channel
	fmt.Println("Output channel", data)
}
