package main

import (
	"fmt"
	"time"
)

const LIM = 45

var fib [LIM]int

func main() {
	result := 0
	start := time.Now()
	for i := 1; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

	f := fibonacci2()
	for i := 2; i < 9; i++ {
		println(i, f())
	}
}

// 有缓存的菲波那契数列
func fibonacci(n int) (res int) {
	if fib[n] != 0 {
		res = fib[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fib[n] = res
	return
}

func fibonacci2() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
