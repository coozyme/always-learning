package main

import (
	"fmt"
	"time"
)

func HelloWord() {
	fmt.Println("Hello World")
}

func DiplayNumber(i int) {
	fmt.Println("display number int", i)
}

func main() {
	start := time.Now()
	go HelloWord()

	fmt.Println("Hai")

	for i := 0; i < 1000; i++ {
		go DiplayNumber(i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("timess", time.Since(start))
}
