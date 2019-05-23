package main

import (
	"fmt"
	"time"
)

func main() {
	result := 0
	start := time.Now()
	for i := 0; i < 41; i++ {
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

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func fibonacci2() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
