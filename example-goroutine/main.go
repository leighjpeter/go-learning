package main

import (
	"fmt"
	"strconv"
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

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
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
func main() {
	// go test()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("main() hello world" + strconv.Itoa(i))
	// 	time.Sleep(time.Second)
	// }
	var mychan chan int
	mychan = make(chan int, 3)
	fmt.Printf("mychan的值是%v，mychan本身的地址%p\n", mychan, &mychan)
	// write data
	// 写入数据不能超过cap
	mychan <- 10
	num := 12
	mychan <- num
	fmt.Printf("mychan的长度是%v，mychan的容量是%v\n", len(mychan), cap(mychan))
	// get data
	// 读取数据不能无限读
	var num_get1 int
	num_get1 = <-mychan
	fmt.Println(num_get1)
	fmt.Printf("mychan的长度是%v，mychan的容量是%v\n", len(mychan), cap(mychan))

	// 循环channel 必须要close()
	c := make(chan bool)
	go func() {
		fmt.Println("GOGOGO")
		c <- true
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}

	return
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
