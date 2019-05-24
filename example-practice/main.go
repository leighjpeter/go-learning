package main

import (
	"fmt"
	"github.com/leighjpeter/go-learning/example-practice/model"
	"math/rand"
	"sort"
	// "strconv"
	"github.com/gomodule/redigo/redis"
	"math"
	"time"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:   8, //最大空闲链接数
		MaxActive: 0, // 表示和数据库的最大链接数， 0 表示没有限制 IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接哪个 ip 的 redis
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

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

// 冒泡排序
func BubbleSort(arr *[5]int) {
	temp := 0
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := 0; j < i-1; j++ {
			if (*arr)[i] > (*arr)[j+1] {
				temp = (*arr)[i]
				(*arr)[i] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
}

func BubbleSort2() {
	//================冒泡=======================//
	unsort_a := [...]int{9, 6, 3, 8, 5}
	l := len(unsort_a)
	fmt.Println(unsort_a)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if unsort_a[i] > unsort_a[j] {
				unsort_a[i], unsort_a[j] = unsort_a[j], unsort_a[i]
			}
		}
	}
	fmt.Println(unsort_a)

}

// 快速排序
func quickSort(arr []int, leftIndex, rightIndex int) {
	if leftIndex > rightIndex {
		return
	}
	i := leftIndex
	j := rightIndex
	base := arr[i]

	for i != j {
		for arr[j] > base && j > i {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			// i++
		}

		for arr[i] <= base && j > i {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			// j--
		}
	}
	arr[i] = base
	quickSort(arr, leftIndex, i-1)
	quickSort(arr, i+1, rightIndex)
}

// 快排方式2
func quickSort2(arr []int, leftIndex, rightIndex int) {
	i, j := leftIndex, rightIndex
	base := arr[i]
	for i < j {
		for j > i && arr[j] > base {
			j--
		}
		if j >= i {
			arr[i] = arr[j]
		}

		for i < j && arr[i] <= base {
			i++
		}
		if i < j {
			arr[j] = arr[i]
		}
	}

	arr[i] = base
	quickSort(arr, leftIndex, i-1)
	quickSort(arr, i+1, rightIndex)
}

// 有序数组的二分查找
func BinaryFind(arr *[6]int, leftIndex, rightIndex, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("not find")
	}

	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Println("find!")
	}
}

type Student struct {
	Name  string
	Age   int
	Score float64
}

type StuSlics []Student

func (ss StuSlics) Len() int {
	return len(ss)
}
func (ss StuSlics) Less(i, j int) bool {
	return ss[i].Age > ss[j].Age
}
func (ss StuSlics) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
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

func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

func main() {
	f := Round(0.3452135, 2)
	fmt.Println(f)

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

	map_a := make(map[string]int)
	mmmm, ok := map_a["aaa"]
	fmt.Println(mmmm, ok)

	// 排序
	var arr = []int{6, 1, 2, 4, 9, 3, 7, 5, 10, 8, 0}
	len2 := len(arr)
	quickSort2(arr, 0, len2-1)
	fmt.Println(arr)

	// 打印金字塔
	// var n int
	// fmt.Print("输入金字塔的层数：")
	// fmt.Scanln(&n)
	// printPyramid(n)

	// 生成随机数
	rand.Seed(time.Now().UnixNano())
	var intArr [3]int
	len1 := len(intArr)
	for i := 0; i < len1; i++ {
		intArr[i] = rand.Intn(100)
		fmt.Println(intArr[i])
	}

	// 工厂模式
	var st = model.NewStudent("leighj")
	fmt.Println(st)
	fmt.Println(st.Name)
	st.SetScore(100)
	fmt.Println(st.GetScore())

	//实现对 Student 结构体切片的排序
	// b := fmt.Sprintf("%.4f", rand.Float64())
	// a, _ := strconv.ParseFloat(b, 64)
	// fmt.Print(a * 100)
	var ss StuSlics
	for i := 0; i < 10; i++ {
		stu := Student{
			Name: fmt.Sprintf("英雄|%d", rand.Intn(100)),
			Age:  rand.Intn(30),
		}
		ss = append(ss, stu)
	}

	fmt.Println(ss)
	sort.Sort(ss)
	for _, v := range ss {
		fmt.Println(v)
	}

	//编写一个函数判断输入参数的类型
	type_a := 1
	type_b := 1.1
	type_c := true
	type_d := "a"
	stu1 := Student{}
	stu2 := &Student{}
	TypeJudge(type_a, type_b, type_c, type_d, stu1, stu2)

	//
	type User struct {
		id  int
		age int
	}
	var m map[string]*User
	m = make(map[string]*User)
	// insert
	m["woofy"] = &User{
		id:  1001,
		age: 28,
	}
	// addressing:  compile error
	// foo(&m["woofy"])
	// modify: compile error
	m["woofy"].id = 10001
	m["woofy"].age++
	fmt.Println(m["woofy"].age)

	// go 连接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()

	fmt.Println("connect success")

	_, err = conn.Do("Set", "name", "leighj")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	// 返回的r是interface{}
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
	}
	// r = r.(string)  panic: interface conversion: interface {} is []uint8, not string
	fmt.Println(r)
	//设置过期时间
	_, err = conn.Do("expire", "name", 10)

	_, err = conn.Do("HSet", "user01", "name", "leighj")
	r, err = redis.String(conn.Do("HGet", "user01", "name"))
	_, err = conn.Do("HSet", "user01", "age", 18)
	i, err := redis.Int(conn.Do("HGet", "user01", "age"))
	//list
	_, err = conn.Do("lpush", "heroList", "no1:宋江", 30, "no2:卢俊义", 28)
	r, err = redis.String(conn.Do("rpop", "heroList"))

	fmt.Println(r)
	fmt.Println(i)
}
