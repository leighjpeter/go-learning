package main

import (
	"fmt"
	"github.com/leighjpeter/go-learning/example-goroutine/gsema"
	"sync"
	"time"
)

func do(num int, wg *sync.WaitGroup) {
	fmt.Printf("start job:%d\n", num)
	time.Sleep(1 * time.Second)
	fmt.Printf("stop job:%d\n", num)
	wg.Done()
}

func cal(ch chan int, a int, b int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	ch <- 1
}

func WriteData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
	}
	close(intChan)
}

func ReadData(intChan chan int, exitChan chan bool) {
	for {
		x, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("read data:%v\n", x)
	}

	exitChan <- true
	close(exitChan)
}

func add(a, b int) {
	c := a + b
	fmt.Printf("%d + %d = %d", a, b, c)
}

func main() {
	go add(1, 2)
	time.Sleep(1e9)
	return
	// data := make(chan int)
	// exit := make(chan bool)

	// go func() {
	// 	for d := range data {
	// 		fmt.Println(d)
	// 	}
	// 	fmt.Println("reci over.")
	// 	exit <- true
	// }()

	// data <- 1
	// data <- 2
	// data <- 3
	// close(data)
	// fmt.Println("send over.")
	// for <-exit {
	// 	return
	// }

	var exitChan chan bool
	var intChan chan int

	exitChan = make(chan bool, 1)
	intChan = make(chan int, 50)

	go WriteData(intChan)
	go ReadData(intChan, exitChan)

	for !<-exitChan {
		break
	}

	chan_a := make(chan int)
	for i := 0; i < 10; i++ {
		go cal(chan_a, i, i+1)
	}
	<-chan_a

	// Wait()
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go do(i, &wg)
	}
	wg.Wait()
	fmt.Println("DONE")

	/*
		//无缓存是同步阻塞的
		chan_a := make(chan int)
		go func() {
			fmt.Println("GOGOGO!!!")
			chan_a <- 1
		}()
		<-chan_a

		chan_c := make(chan int)
		go func() {
			fmt.Println("GOGOGO!!!")
			<-chan_c
		}()
		chan_c <- 1

		//有缓存是异步的，放先于取
		chan_b := make(chan int, 1)
		go func() {
			fmt.Println("GOGOGO!!!")
			chan_b <- 1
		}()
		<-chan_b
	*/
}

// 控制 goroutine的并发数量
var wg = sync.WaitGroup{}

func main2() {
	userCount := 5
	ch := make(chan bool, 3)
	for i := 0; i < userCount; i++ {
		wg.Add(1)
		go Read(ch, i)
	}
	wg.Wait()
}

func Read(ch chan bool, i int) {
	defer wg.Done()
	ch <- true
	fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
	time.Sleep(time.Second)
	<-ch
}

// 控制goroutine2个并发
var sema = gsema.NewSemaphore(2)

func main3() {
	userCount := 5
	for i := 0; i < userCount; i++ {
		go ReadSema(i)
	}
	sema.Wait()
	time.Sleep(2 * time.Second)
}

func ReadSema(i int) {
	defer sema.Done()
	sema.Add(1)
	fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
	time.Sleep(time.Second)
}

// 变更channel的输入数量
// 变更channel的循环值
// 变更最大允许并发的goroutine数量
func main4() {
	userCount := 10
	ch := make(chan int, 5)
	for i := 0; i < userCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for d := range ch {
				fmt.Printf("go func: %d, time: %d\n", d, time.Now().Unix())
				time.Sleep(time.Second * time.Duration(d))
			}
		}()
	}

	for i := 0; i < 10; i++ {
		ch <- 1
		ch <- 2
		//time.Sleep(time.Second)
	}

	close(ch)
	wg.Wait()
}
