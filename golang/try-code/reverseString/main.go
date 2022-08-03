package main

import "fmt"

func reverseString(val string) {
	// valArr := []byte{}

	for _, v := range val {
		fmt.Println(v)
	}

}

func main() {
	reverseString("oke")
}
