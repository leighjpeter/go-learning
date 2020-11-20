package main

/**
* 协程间基于共享内存的方式进行通信
* 缺点：为了保证操作的原子性，引入互斥锁，加锁解锁麻烦，代码复杂。
 */
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int = 0

func add(a, b int, lock *sync.Mutex) {
	c := a + b
	lock.Lock()
	counter++
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	defer lock.Unlock()
}

func main() {
	start := time.Now()

	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go add(1, i, lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	fmt.Printf("consume: %s\n", time.Since(start))
}
