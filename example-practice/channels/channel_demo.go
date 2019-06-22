package main

import (
	"fmt"
	"time"
)

func main() {
	//速率限制
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

	//定时器和打点器
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
