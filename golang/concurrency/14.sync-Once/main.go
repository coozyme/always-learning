package main

import (
	"fmt"
	"sync"
)

var count int = 0

func DoOnce() {
	count++
	fmt.Println(count)
}
func main() {
	var group sync.WaitGroup
	var once sync.Once

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			once.Do(DoOnce)
			group.Done()
		}()
	}
	group.Wait()

	fmt.Println("DONES")
}
