package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadFile() (string, error) {
	fileName := "read.txt"

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
		return "", err
	}

	return string(data), nil
}

func main() {
	readF, _ := ReadFile()
	fmt.Println("re", string(readF))
}
