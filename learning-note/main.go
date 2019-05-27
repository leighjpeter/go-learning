package main

import (
	"fmt"
	"math/rand"
	_ "runtime"
	// "sort"
	_ "sync"
	"time"
	"unicode/utf8"
)

type File struct {
	name string
	size int
	attr struct {
		perm  int
		owner int
	}
}

type Tester interface {
	Do()
}
type FuncDo func()

func (self FuncDo) Do() { self() }

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}

func main() {
	f := File{
		name: "test.txt",
		size: 1024,
	}

	f.attr.perm = 0755
	f.attr.owner = 1

	fmt.Println(f)

	var t Tester = FuncDo(func() {
		println("Hello, World!")
	})
	t.Do()

	rand.Seed(time.Now().UnixNano())
	println(time.Second)

	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))                       //28
	fmt.Printf("The number of characters in string str2 is %d", utf8.RuneCountInString(str2)) //24

	var barVal = map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
	println()
	// time.Sleep(1e9)
	f1() // 4,3,2,1,0
	f2() // 5,5,5,5,5
	f3() // 4,3,2,1,0
	// println(f4()) //5
	// println(f5()) //1
	// println(f6()) //5
	// println(f7()) //1
}

func f1() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func f2() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func f3() {
	for i := 0; i < 5; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func f4() int {
	t := 5
	defer func() {
		t++
	}()
	return t
}

func f5() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f6() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f7() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
