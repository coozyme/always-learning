package main

import (
	"fmt"
	"time"
)

/*
	Buffer capacity
	- kita bebas memasuka jumlah capacity antrian didalam buffer
	- jika kita set 5, artinya kita bisa menerima 5data di buffer
	  dan jika kita masukan 6 data pada buffer, maka kita disuruh nunggu sampai buffer ada yang kosong
	- ini cocok sekali jika menerima data lebih lambat dari pada pengirim

*/
func main() {

	channel := make(chan string, 5) // by default buffer bernilai 1

	defer close(channel)

	go func() {
		channel <- "MASUKAN1"
		channel <- "MASUKAN2"
		channel <- "MASUKAN3"

	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(cap(channel))
	fmt.Println(len(channel))
}
