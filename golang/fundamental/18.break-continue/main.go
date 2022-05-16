package main

import "fmt"

func main() {

	funcContinue()
	funcBreak()
	funcBreakContinue()

}

func funcContinue() {
	// contoh continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("loop continue ke", i)
	}
}

func funcBreak() {
	// contoh break
	for i := 0; i < 10; i++ {
		fmt.Println("loop break ke", i)
		if i == 5 {
			break
		}
	}
}

func funcBreakContinue() {
	// contoh gabungan break dan continue
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		}

		if i == 15 {
			break
		}

		fmt.Println("loop ke", i)
	}
}
