package main

import (
	"fmt"
	"sync"
	"time"
)

type BankBallance struct {
	mutex    sync.RWMutex
	ballance int
}

func main() {

	account := BankBallance{}
	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				account.AddBallance(1)
				fmt.Println(account.GetBallance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total Ballance = ", account.GetBallance())
}

func (b *BankBallance) AddBallance(amount int) {
	b.mutex.Lock()
	b.ballance += amount
	b.mutex.Unlock()
}

func (b *BankBallance) GetBallance() int {
	b.mutex.RLock()
	ballance := b.ballance
	b.mutex.RUnlock()

	return ballance

}
