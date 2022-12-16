package main

// 1
import (
	"fmt"

	"gopkg.in/robfig/cron.v2"
)

func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}

func runCronJobs() {
	// 2
	s := cron.New()

	// 3
	s.AddFunc("@every 1s", func() {
		hello("John Doe")
	})

	// 4
	s.Start()
}

// 5
func main() {
	runCronJobs()
	fmt.Scanln()
}
