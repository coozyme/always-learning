package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	x := 0

	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("X = ", x)
}
