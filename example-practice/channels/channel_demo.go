package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

func main() {
	cc := make(chan int, 1)
	for i := 0; i < 5; i++ {
		cc <- i
	}
	close(cc)
	for v := range cc {
		fmt.Printf("Received:%d\n", v)
	}
	// fmt.Println(<-c)
	return

	// 通道并发安全。词法安全
	// 在chanOwner的词法范围内实例化通道，使得通道的写入操作范围被限制在它定义的闭包中。
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			for i := 0; i < 5; i++ {
				results <- i
			}
			defer close(results)
		}()
		return results
	}
	// 将通道的用法限制为只读
	consumer := func(results <-chan int) {
		for v := range results {
			fmt.Printf("Received:%d\n", v)
		}
		fmt.Println("Done receiving!")
	}
	results := chanOwner()
	consumer(results)
	return

	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}
	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	wg.Wait()
	// return

	// 速率限制
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)
	limter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		<-limter
		fmt.Println("request", req, time.Now())
	}

	// 定时器和打点器
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("Timer 1 expired")
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")

}
