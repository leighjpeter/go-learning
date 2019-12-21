package main

import (
	"fmt"
	"github.com/leighjpeter/go-learning/example-practice/model"
	"math/rand"
	"sort"
	// "strconv"
	"bufio"

	"math"
	"net"
	"os"
	"sync"
	"time"
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

// 冒泡排序
func BubbleSort2() {
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

// 选择排序
func SelectSort(arr []int) {
	n := len(arr)
	for j := 0; j < n-1; j++ {
		max := arr[j]
		maxIndex := j
		for i := j + 1; i < n; i++ {
			if max <= arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}
}

// 插入排序
func InsertSort(a []int) {
	n := len(a)
	for i := 1; i < n; i++ { //如果第i个元素大于第j个，则第j个向后移动
		v := a[i]
		j := i - 1
		for ; j >= 0 && v > a[j]; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = v
	}
}

// 快速排序
func QuickSort(arr []int, leftIndex, rightIndex int) {
	if leftIndex > rightIndex {
		return
	}
	i, j := leftIndex, rightIndex
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
	QuickSort(arr, leftIndex, i-1)
	QuickSort(arr, i+1, rightIndex)
}

// 快排方式2
func QuickSort2(arr []int, leftIndex, rightIndex int) {
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
	QuickSort(arr, leftIndex, i-1)
	QuickSort(arr, i+1, rightIndex)
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

func hannuota(n int, A, B, C string) {
	if n < 1 || n > 20 {
		return
	}
	if n == 1 {
		fmt.Printf("盘子%d从%s柱子移动到%s柱子\n", n, A, C)
	} else {
		hannuota(n-1, A, C, B)
		fmt.Printf("盘子%d从%s柱子移动到%s柱子\n", n, A, C)
		hannuota(n-1, B, A, C)
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

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teachershowB")
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
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

func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		},
		func() {
			println(x)
		}
}
func main() {
	var sli = []int{6, 1, 2, 4, 9, 3, 7, 5, 10, 8, 0}
	// len2 := len(sli)
	// QuickSort2(sli, 0, len2-1)
	// fmt.Println(sli)

	// SelectSort(sli)
	InsertSort(sli)
	fmt.Println(sli)
	return
	var slice []int
	slice = make([]int, 0)
	println(slice == nil)
	return
	a1, b1 := test(100)
	a1()
	b1()
	return
	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	vvv := <-th.lter()

	fmt.Sprintf("%s%v", "ch", vvv)

	return
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
	return

	t := Teacher{}
	t.ShowA()
	t.ShowB()
	return
	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		addr := net.ParseIP(sender.Text())
		if addr == nil {
			fmt.Println("Invalid address")
			os.Exit(0)
		} else {
			fmt.Println("The address is ", addr.String())
		}
	}

	// 排序
	// var sli = []int{6, 1, 2, 4, 9, 3, 7, 5, 10, 8, 0}
	// len2 := len(sli)
	// QuickSort2(sli, 0, len2-1)
	// fmt.Println(sli)

	// SelectSort(sli)
	// InsertSort(sli)
	// fmt.Println(sli)
	// 汉诺塔
	hannuota(3, "A", "B", "C")
	return
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
	// a := [...]int{1, 2, 3, 4, 5}
	// for i, v := range a {
	// 	a[2] = 100
	// 	if i == 2 {
	// 		fmt.Println(i, v) // 2, 3
	// 	}
	// }
	// fmt.Println(a) // [1 2 100 4 5 ]

	n := [...]int{1, 2, 3, 4, 5}
	for i, v := range &n {
		n[2] = 100
		if i == 2 {
			fmt.Println(i, v) // 2 100
		}
	}
	fmt.Println(n) // [1 2 100 4 5 ]
	// return

	map_a := make(map[string]int)
	mmmm, ok := map_a["aaa"]
	fmt.Println(mmmm, ok)

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
	// var st = model.NewStudent("leighj")
	// fmt.Println(st)
	// fmt.Println(st.Name)
	// st.SetScore(100)
	// fmt.Println(st.GetScore())

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

}
