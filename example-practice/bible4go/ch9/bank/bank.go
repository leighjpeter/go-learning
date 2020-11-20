package bank

// 当你使用mutex时，确保mutex和其保护的变量没有被导出(在go里也就是小写，且不要被大写字母开头的函数访问啦)，无论这些变量是包级的变量还是一个struct的字段。

import (
	"sync"
)

var (
	mu      sync.Mutex // guards balance
	balance int
)

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func deposit(amount int) {
	balance += amount
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if Balance() < 0 {
		deposit(amount)
		return false
	}
	return true
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}

	}
}

func init() {
	go teller()
}
