package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	pull := sync.Pool{}

	pull.Put("Ary")
	pull.Put("Setya")
	pull.Put("Pambudi")

	for i := 0; i < 10; i++ {
		go func() {
			data := pull.Get()

			fmt.Println(data)
			pull.Put(data)
		}()
	}
	time.Sleep(2 * time.Second)

	fmt.Println("DONE")
}
