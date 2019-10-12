package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

func main() {
	// 通道并发安全。词法安全
	// 在chanOwner的词法范围内实例化通道，使得通道的写入操作范围被限制在它定义的闭包中。
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				results <- i
			}
		}()
		return results
	}
	// 将通道的用法限制为只读
	constomer := func(results <-chan int) {
		for v := range results {
			fmt.Printf("Received:%d\n", v)
		}
		fmt.Println("Done reciving!")
	}

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
	return
	results := chanOwner()
	constomer(results)

	return

	// 超时控制
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("time out")
	}

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

	//通道关闭
	/*
		jobs := make(chan int, 5)
		done := make(chan bool)
		go func() {
			for {
				j, ok := <-jobs
				if ok {
					fmt.Println("received job:", j)
				} else {
					fmt.Println("received all jobs")
					done <- true
					return
				}
			}
		}()

		for i := 0; i < 3; i++ {
			jobs <- i
			fmt.Println("sent job", i)
		}
		close(jobs)
		fmt.Println("sent all jobs")
		<-done
	*/
	//非阻塞通道操作
	/*
		messages := make(chan string)
		signals := make(chan bool)

		select {
		case msg := <-messages:
			fmt.Println("received message:", msg)
		default:
			fmt.Println("no message received")
		}

		msg := "hi"
		select {
		case messages <- msg:
			fmt.Println("sent message", msg)
		default:
			fmt.Println("no message sent")
		}

		select {
		case msg := <-messages:
			fmt.Println("received message", msg)
		case sig := <-signals:
			fmt.Println("received signal", sig)
		default:
			fmt.Println("no activity")
		}
	*/
}
