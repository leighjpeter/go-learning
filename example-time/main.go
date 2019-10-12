package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("PRC")
	second, _ := time.ParseDuration("10h")
	fmt.Println(second.Seconds())

	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	tt, _ = time.Parse("2006-01-02 15:04:05", timeStr)

	fmt.Println(tt.Location())
	//如果不设置loc， Unix返回早八点的时间戳，减去8个小时
	timestamp := tt.UTC().Unix()
	fmt.Println("timestamp:", timestamp)
	str := "2019-07-17 23:59:59"
	t, _ := time.Parse("2006-01-02 15:04:05", str)
	println(t.Unix() - 8*3600)

	now := time.Now()
	fmt.Printf("now=%v type=%T\n", now, now)
	//now=2019-07-17 17:02:58.854298 +0800 CST m=+0.000364769 type=time.Time
	fmt.Printf("unix=%v unixnano=%v\n", now.Unix(), now.UnixNano())
	// unix=1563354178 unixnano=1563354178854298000

	// Timer
	fmt.Println(time.Now())
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println(time.Now())

	// STOP
	fmt.Println("start")
	timer := time.NewTimer(time.Second * 2)
	go func() {
		<-timer.C
		fmt.Println("get timer")
	}()
	if timer.Stop() {
		fmt.Println("timer stoped")
	}

	//AfterFunc
	wait := sync.WaitGroup{}
	fmt.Println("start:", time.Now())
	wait.Add(1)
	timer2 := time.AfterFunc(time.Second*3, func() {
		fmt.Println("get timer", time.Now())
		wait.Done()
	})
	time.Sleep(time.Second)
	fmt.Println("sleep", time.Now())
	timer2.Reset(time.Second * 4)
	wait.Wait()

	// Ticker
	fmt.Println("start ticker", time.Now())
	ticker := time.NewTicker(time.Second)
	go func() {
		for tick := range ticker.C {
			fmt.Println("tick at", tick)
		}
	}()
	time.Sleep(time.Second * 5)
	ticker.Stop()
	fmt.Println("stoped", time.Now())
	return
	// buildin函数(内置函数)
	// new 一般创建值类型
	// make 一般创建引用类型
	num := new(int)
	fmt.Printf("类型=%T 值=%v 地址=%v 指针指向的值%v \n", num, num, &num, *num)

	var numArr01 [3]int = [3]int{1, 2, 9223372036854775807}
	var numArr02 = [3]int{1, 2, 3}
	var numArr03 = [...]int{1, 2, 3}
	var numArr04 = [...]int{1: 800, 0: 900, 2: 999}
	strArr05 := [...]string{"a", "b", "c"}
	fmt.Println(numArr01, numArr02, numArr03, numArr04, strArr05)

	// 切片动态追加
	var slice_a []int = make([]int, 2)
	fmt.Println(slice_a)

	slice_a = append(slice_a, slice_a...)
	fmt.Println(slice_a)

	// 切片拷贝
	var slice_1 []int = []int{1, 2, 3}
	var slice_2 = make([]int, 10)

	copy(slice_2, slice_1)

	fmt.Println(slice_1)
	fmt.Println(slice_2)

}
