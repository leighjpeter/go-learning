package bank_test

import (
	"fmt"
	"testing"

	"github.com/leighjpeter/go-learning/example-practice/bible4go/ch9/bank"
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

	go func() {
		bank.Withdraw(100)
		done <- struct{}{}
	}()

	<-done
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

}
