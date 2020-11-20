// 编写一个程序，开启 3 个线程A,B,C，这三个线程的输出分别为 A、B、C，
// 每个线程将自己的 输出在屏幕上打印 10 遍，
// 要求输出的结果必须按顺序显示。如：ABCABCABC....

package main

import (
	"fmt"
	"log"
	// "log"
)

var (
	V2N = 3
	V2M = 5
)

func echo(threadNum int, wait, sig chan struct{}) {
	threadName := string('A' + threadNum)
	for range wait {
		fmt.Printf("%d: %s\n", threadNum, threadName)
		sig <- struct{}{}
	}
	close(sig)
}

// 利用channel做信号量的解法
// 考察同步实现
func main() {

	var wait, sig, firstWait, lastSig chan struct{}

	wait = make(chan struct{})
	firstWait = wait

	for i := 0; i < V2N; i++ {
		sig = make(chan struct{})
		lastSig = sig
		go echo(i, wait, sig)
		wait = sig
	}

	for i := 0; i < V2M; i++ {
		firstWait <- struct{}{}
		<-lastSig
	}
	close(firstWait)

	if _, ok := <-lastSig; ok {
		log.Fatalln("Channel not closed")
	}

}
