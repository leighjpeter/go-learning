package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("now=%v type=%T\n", now, now)
	// 格式化日期时间
	// 2006-01-02 15:04:05 数字必须固定
	s := now.Format("2006-01-02 15:04:05")
	fmt.Println(s)
	// unix unixNano
	fmt.Printf("unix=%v unixnano=%v\n", now.Unix(), now.UnixNano())

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
