package main

import (
	"fmt"
	"math"
	// "math/rand"
	// "net"
	// "os"
	// "runtime"
	// "sort"
	"sync"
	// "time"
)

// 金字塔
func printPyramid(n int) {
	// 打印层数
	for i := 1; i <= n; i++ {
		// 打印空格
		for k := 0; k < n-i; k++ {
			fmt.Print(" ")
		}
		// 打印符号
		for j := 0; j < (2*i - 1); j++ {
			if j == 0 || j == (2*i-2) || i == n {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

type Student struct {
	Name  string
	Age   int
	Score float64
}

// 判断参数类型
func TypeJudge(items ...interface{}) {
	for i, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第%v的参数类型是bool,值是%v\n", i, v)
		case float64, float32:
			fmt.Printf("第%v的参数类型是float,值是%v\n", i, v)
		case int, int32, int64:
			fmt.Printf("第%v的参数类型是int,值是%v\n", i, v)
		case string:
			fmt.Printf("第%v的参数类型是string,值是%v\n", i, v)
		case Student:
			fmt.Printf("第%v的参数类型是Student,值是%v\n", i, v)
		case *Student:
			fmt.Printf("第%v的参数类型是指针Student,值是%v\n", i, v)
		default:
			fmt.Printf("第%v的参数类型不明确,值是%v\n", i, v)
		}
	}
}

// 保留小数位
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) lter() <-chan interface{} {
	ch := make(chan interface{}, len(set.s))
	// ch := make(chan interface{})
	go func() {
		set.RLock()
		for elem, value := range set.s {
			ch <- elem
			fmt.Println("Iter:", elem, value)
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

func main() {

	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	vvv := <-th.lter()

	fmt.Sprintf("%s%v", "ch", vvv)

	return

	v := [...]int{1, 2, 3, 4, 5}
	for i := range v {
		v[2] = 100
		if i == 2 {
			fmt.Println(i, v[i]) //2,100
		}
	}
	fmt.Println(v) // [1 2 100 4 5]

	a := [...]int{1, 2, 3, 4, 5}
	for i, v := range a {
		a[2] = 100
		if i == 2 {
			fmt.Println(i, v) // 2, 3
		}
	}
	fmt.Println(a) // [1 2 100 4 5 ]

	n := [...]int{1, 2, 3, 4, 5}
	for i, v := range &n {
		n[2] = 100
		if i == 2 {
			fmt.Println(i, v) // 2 100
		}
	}
	fmt.Println(n) // [1 2 100 4 5 ]
	return

	// 打印金字塔
	// var n int
	// fmt.Print("输入金字塔的层数：")
	// fmt.Scanln(&n)
	// printPyramid(n)

}
