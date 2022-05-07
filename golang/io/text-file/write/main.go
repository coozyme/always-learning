package main

import (
	"fmt"
	"log"
	"os"
)

func WriteFile() {
	// Create file
	file, err := os.Create("write.txt")

	if err != nil {
		log.Fatalf("failed create file, err: %s", err)
	}

	defer file.Close()

	len, err := file.WriteString("Write anything here\n" +
		"This program demonstrates reading and writing\n" +
		"operations to a file in Go")

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("fn %s \n", file.Name())
	fmt.Println("fl", len)

}

func main() {
	WriteFile()
}
