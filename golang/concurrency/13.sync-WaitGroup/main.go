package main

import (
	"fmt"
	"sync"
	"time"
)

func RunAsyc(i int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Printf("%d running async \n", i)
	time.Sleep(1 * time.Second)
}
func main() {
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go RunAsyc(i, group)
	}

	group.Wait()

	fmt.Println("DONES")
}
