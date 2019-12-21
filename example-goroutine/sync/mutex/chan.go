package main

/**
* 协程间基于channel进行通信
 */
import (
	"fmt"
	"time"
)

var counter int = 0

func add(a, b int, ch chan int) {
	c := a + b
	counter++
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	ch <- 1
	close(ch)
}

func main() {
	start := time.Now()

	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go add(1, i, chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("consume:", consume)
}
