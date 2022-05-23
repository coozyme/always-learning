package main

import "fmt"

func endApp() {
	message := recover()

	if message != nil {
		fmt.Println("Error message: ", message)
	}

	fmt.Println("End App")
}

func runApp(err bool) {

	if err {
		defer endApp()
		panic("Terjadi error")
	}

	fmt.Println("App Running")

}
func main() {
	// runApp(true)
	runApp(false)

}
