// 编写一个程序，开启 3 个线程A,B,C，这三个线程的输出分别为 A、B、C，
// 每个线程将自己的 输出在屏幕上打印 10 遍，
// 要求输出的结果必须按顺序显示。如：ABCABCABC....

package main

import (
	"fmt"
)

var (
	N = 5
	M = 5
)

func gen(s string, times int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < times; i++ {
			ch <- s
		}
	}()
	return ch
}

func fanIn(times int, inputs []<-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < times; i++ {
			for _, input := range inputs {
				v := <-input
				ch <- v
			}
		}
	}()
	return ch
}

func main() {
	times := M
	i := 1
	inputs := make([]<-chan string, 0, N)
	for i := 0; i < N; i++ {
		threadName := string('A' + i)
		inputs = append(inputs, gen(threadName, times))
	}

	for char := range fanIn(M, inputs) {
		fmt.Printf("%d: %s\n", i, char)
		i++
	}

}
