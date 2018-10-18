package bank

import (
	"fmt"
	"sync"
	"testing"

	"./bank"
)

func TestBank(t *testing.T) {
	var wg sync.WaitGroup

	// Alice
	wg.Add(1)
	go func() {
		defer wg.Done()
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
	}()

	// Bob
	wg.Add(1)
	go func() {
		defer wg.Done()
		bank.Withdraw(100)
	}()

	wg.Wait()

	if got, want := bank.Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
