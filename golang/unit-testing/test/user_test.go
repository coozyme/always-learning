package main

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	user1 := User{Name: "asd", Age: 123}

	user1.Age = 10
	fmt.Println(user1)

	New := NewUser("Ary", 12)

	New.Age = 20

	fmt.Println(New)
}
