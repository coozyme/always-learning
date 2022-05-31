package main

import (
	"fmt"
	"sync"
)

/*
	Schema
	melakukan transfer secara bersamaan,
	User1   ->	user2
	User2   ->	user1

*/

type Account struct {
	sync.Mutex
	Name     string
	Ballance int
}

func main() {
	ary := Account{
		Name:     "Ary",
		Ballance: 10000,
	}

	setya := Account{
		Name:     "Setya",
		Ballance: 15000,
	}

	go Transfer(&ary, &setya, 2000)
	go Transfer(&setya, &ary, 5000)

	// time.Sleep(3 * time.Second)

	fmt.Println("user-1 : ", ary.Name, "ballance : ", ary.Ballance)
	fmt.Println("user-2 : ", setya.Name, "ballance : ", setya.Ballance)
}

func Transfer(user1, user2 *Account, amount int) {
	user1.Mutex.Lock()
	fmt.Println("Lock user 1", user1.Name)
	user1.ChangeBallance(-amount)

	user2.Mutex.Lock()
	fmt.Println("Lock user 2", user2.Name)
	user2.ChangeBallance(amount)

	user1.Mutex.Unlock()
	user2.Mutex.Unlock()

}

func (user *Account) ChangeBallance(amount int) {
	user.Ballance = user.Ballance + amount
}
