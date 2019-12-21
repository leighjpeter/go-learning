package demo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGor00(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestGor02(t *testing.T) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Print("panic!!!,recover")
			}
		}()
		panic("...")
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}
