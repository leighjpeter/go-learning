package main

/**
* 协程间基于channel进行通信
 */
import (
	"fmt"
	"sync"
	"time"
)

var counter int = 0

func add(a, b int, wg *sync.WaitGroup) {
	c := a + b
	// counter++
	fmt.Printf(" %d + %d = %d\n", a, b, c)
	wg.Done()
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go add(1, i, &wg)
	}
	wg.Wait()
	fmt.Printf("consume: %s\n", time.Since(start))
}
