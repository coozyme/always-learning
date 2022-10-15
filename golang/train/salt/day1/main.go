package main

import (
	"always-learning/golang/train/salt/day1/entity"
	"always-learning/golang/train/salt/day1/valueobject"
	"fmt"
)

func main() {

	ary := entity.NewUser(
		valueobject.NewName("ary", "setya"),
		valueobject.NewEmail("ary", "salt.co.id"),
		valueobject.NewPassword("salt", "hellow"),
	)

	fmt.Println(ary.Name.Full())
	fmt.Println(ary.Email.ToString())
	fmt.Println(ary.Password.Verify("hellow"))
}
