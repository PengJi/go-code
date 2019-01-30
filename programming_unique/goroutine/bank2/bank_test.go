package bank_test

import (
	"sync"
	"testing"

	"github.com/pengji/go-code/programming_unique/goroutine/bank2"
)

func TestBank(t *testing.T) {
	// Deposit [1..1000] concurrently.
	var n sync.WaitGroup
	for i := 1; i<= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank.Deposit(amount)
			n.Done()
		}(i)
	}
}