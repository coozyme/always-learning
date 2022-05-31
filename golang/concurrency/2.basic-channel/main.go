package main

import (
	"fmt"
	"time"
)

func main() {
	// create channel
	channel := make(chan string)
	defer close(channel)

	/*
		cara channel Menerima data
		channel <- data


		Jika channel tidak ada yang nerima / mengirim data pada channel tersebut maka aplikasi akan deadlock / berhenti
		pastikan channel menerima data atau mengirim data
	*/

	go func() {
		time.Sleep(2 * time.Second)

		channel <- "ary"
		fmt.Println("Selesai mengirim data ke channel")

	}()
	data := channel

	// fmt.Println(<- channel)
	fmt.Print(data)
	time.Sleep(5 * time.Second)

}
