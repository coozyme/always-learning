package main

import (
	"fmt"
	"strings"
)

func Palindrome(s string) bool {

	words := []byte(s)

	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}

	if strings.EqualFold(s, string(words)) || strings.EqualFold(strings.Replace(s, " ", "", -1), strings.Replace(string(words), " ", "", -1)) {
		fmt.Println("ok", string(words))
		return true
	} else {
		fmt.Println("no ok", string(words))
		return false
	}
}
