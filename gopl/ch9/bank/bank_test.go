package bank_test

import (
	"fmt"
	"go-practice/gopl/ch9/bank"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

}
