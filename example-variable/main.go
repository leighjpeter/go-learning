package main

import (
	"fmt"

	"sync"
	// "sort"
	"encoding/json"
	"reflect"
	"strings"
)

const const_a int = 1
const const_b = 'A'
const const_c, const_d, const_e = 1, 2, 3

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

func A() {
	fmt.Println("func A")
}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("Panic in B")
}

func C() {
	fmt.Println("func C")
}

//自定义数据类型
type TZ int

func (func_a *TZ) Print() {
	fmt.Println("OK")
}
func (func_a *TZ) Increase(num int) {
	*func_a += TZ(num)
}

type TT struct {
	Name string
}

func (func_a *TT) Print(name string) {
	func_a.Name = name
	fmt.Println(func_a.Name)
}

//================Interface=======================//
type USB interface {
	Name() string
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected:", pc.name)
}

func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknow decive")
	}
}

//================Func=======================
func getSum(n1, n2 int) int {
	return n1 + n2
}

func myfun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// 变参函数
func sum(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

//eg:编写一个函数swap 交换n1，n2的值
func swap(n1, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func f7() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
func main() {

	var activityConfig = map[string]float64{
		"0|10":  50,
		"10|50": 75,
		"50|+":  100,
	}

	for k, v := range activityConfig {
		println(k, v)
	}
	return

	type Student struct {
		Name string
		Age  int
	}
	var st3 = Student{"leighj", 30}
	fmt.Println(st3)
	var st1 = Student{Name: "leighj", Age: 30}
	fmt.Println(st1)
	// 返回结构体的指针类型
	var st2 = &Student{Name: "leighj", Age: 30}
	fmt.Println(*st2)

	var closure [2]func()

	for i := 0; i < 2; i++ {
		closure[i] = func() {
			println(i)
		}
	}
	closure[0]() // 2
	closure[1]() // 2

	for i := 0; i < 2; i++ {
		val := i
		closure[i] = func() {
			println(val)
		}
	}
	closure[0]() // 0
	closure[1]() // 1

	type Foo struct {
		bar string
	}
	list := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	list2 := make([]*Foo, len(list))
	for i, value := range list {
		list2[i] = &value
	}
	fmt.Println(list[0], list[1], list[2])    // {A} {B} {C}
	fmt.Println(list2[0], list2[1], list2[2]) // &{C} &{C} &{C}

	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	return
	// Go 语言为数不多的陷阱
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(len(arr), cap(arr)) // 5 5
	slice := arr[1:2]
	fmt.Println(slice, len(slice), cap(slice)) // [2] 1 4
	slice = append(slice, 6, 7, 8)
	fmt.Println(slice)                  // [2,6,7,8]
	fmt.Println(len(slice), cap(slice)) // 4,4
	fmt.Println(arr)                    // [1,2,6,7,8]
	// 信条: 不对函数slice类型的参数append

	return
	ss := []int{5, 6, 7, 8, 9}
	copy(ss[2:], ss[3:])
	fmt.Println(ss)
	return
	//===============String======================
	addBmp := MakeAddSuffix(".bmp")
	s := addBmp("file")
	fmt.Println(s)
	//================Func=======================
	/*
		func_a := getSum(10, 20)
		fmt.Printf("a的类型%T,getSum的类型%T\n", func_a, getSum)
		fmt.Println(func_a)
		func_b := myfun(getSum, 10, 10)
		fmt.Println(func_b)

		func_c := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
		fmt.Println(func_c)

		n1 := 10
		n2 := 20
		swap(&n1, &n2)
		fmt.Printf("n1的值%v，n2的值%v\n", n1, n2)

		//匿名函数
		niming := func(nm1, nm2 int) int {
			return nm1 + nm2
		}(1, 2)
		fmt.Printf("niming的值%v\n", niming)
		return
	*/
	//================CONTINUE/GOTO/BREAK/LABEL=======================//
	/*
		LABEL:
			for L_a := 0; L_a < 3; L_a++ {
				for {
					fmt.Println(L_a)
					continue LABEL
				}
			}

			for {
				for L_a := 0; L_a < 3; L_a++ {
					fmt.Println(L_a)
				}
				goto LABEL1
			}
		LABEL1:
	*/

	//================SLICE=======================//
	/*
		slice_a := make([]int, 0)
		fmt.Print(slice_a)
			a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			fmt.Printf("%p", a)
			slice_a := a[:]
			fmt.Println(slice_a)
			fmt.Printf("%p", slice_a)
	*/

	//================MAP=======================//
	/*
		// 对map间接排序
		map_a := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
		slice_a := make([]int, len(map_a))
		i := 0
		for k, _ := range map_a {
			slice_a[i] = k
			i++
		}
		sort.Ints(slice_a)
		fmt.Println(slice_a)

		// 转换map的key-value
		map_1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
		map_2 := make(map[string]int)

		for k, v := range map_1 {
			map_2[v] = k
		}
		fmt.Println(map_2)
	*/
	//================DEFER=======================//
	// A()
	// B()
	// C()

	//================Struct=======================//
	/*
		type human struct {
			Sex int
		}
		type teacher struct {
			human
			Name string
			Age  int
		}
		type student struct {
			human
			Name string
			Age  int
		}
		//创建方式1，2
		teacher_a := teacher{Name: "joe", Age: 19, human: human{Sex: 0}}
		student_a := student{Name: "joe", Age: 19, human: human{Sex: 1}}
		teacher_a.Name = "leighj"
		student_a.Sex = 2
		teacher_a.human.Sex = 4
		fmt.Println(teacher_a.human)
		fmt.Println(student_a)

		//创建方式3
		var p1 *human = new(human) // p1是个指针 等价于 p1 := new(human)
		p1.Sex = 10
		(*p1).Sex = 20
		fmt.Println(p1)
		//创建方式4
		var p2 *human = &human{} // p2是个指针 等价于 p2 := &human{}
		p2.Sex = 100
		(*p2).Sex = 200
		fmt.Println(p2)

		type struct_C struct {
			Name string
		}

		type struct_B struct {
			struct_C
			Age int
		}

		type struct_A struct {
			struct_B
		}

		b := struct_B{struct_C{"C"}, 10}
		fmt.Println(b.Name, b.Age)
		a := struct_A{struct_B: struct_B{Age: 10, struct_C: struct_C{Name: "C"}}}
		fmt.Println(a.struct_C.Name, a.Age)
	*/

	//================Method=======================//
	/*
		var method_a TZ
		method_a.Print()
		(*TZ).Print(&method_a)

		var method_b TZ
		method_b.Increase(100)
		fmt.Println(method_b)

		st_a := &TT{}
		fmt.Printf("st_a的类型%T\n", st_a)
		st_a.Print("leighj")
		fmt.Println(st_a.Name)
	*/

	//================Interface==================//
	/*
		interface_a := PhoneConnecter{"HUAWEI"}
		interface_a.Connect()
		Disconnect(interface_a)

		var n float64 = 8.8
		var nil_a interface{}
		nil_a = n
		fmt.Print(nil_a)
	*/

	//================JSON=======================//
	/*
		// 序列化
		JsonStruct()
		JsonMap()
		JsonSlice()
		// 反序列化
		UnMarshalStruct()
		UnMarshalMap()
		UnMarshalSlice()
	*/

	//================Reflect==================//
	/*
		//1. 先定义一个 int //
		var num int = 100
		reflectTest01(num)
	*/

}

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func JsonStruct() {
	monster := Monster{Name: "iphone xr", Age: 10, Birthday: "2019-02-01"}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	fmt.Printf("struct序列化后=%v\n", string(data))
}

func JsonMap() {
	var map_a map[string]interface{}
	map_a = make(map[string]interface{})
	map_a["name"] = "leighj"
	map_a["age"] = 11
	data, err := json.Marshal(map_a)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	fmt.Printf("map序列化后=%v\n", string(data))
}

func JsonSlice() {
	var slice_a []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "leighj"
	m1["age"] = 12

	m2 := make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 20

	slice_a = append(slice_a, m1)
	slice_a = append(slice_a, m2)
	data, err := json.Marshal(slice_a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("slice序列化后=%v\n", string(data))
}

func UnMarshalStruct() {

	var monster Monster // monster := Monster{}

	str := "{\"Name\":\"iphone xs\",\"Age\":10,\"Birthday\":\"2019-02-01\",\"Sal\":0,\"Skill\":\"\"}"
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("struct反序列化后=%v,monster.Name=%v\n", monster, monster.Name)
}

func UnMarshalMap() {

	var map_a map[string]interface{}
	//注意:反序列化 map,不需要 make,因为 make 操作被封装到 Unmarshal 函数
	str := "{\"age\":11,\"name\":\"leighj\"}"
	err := json.Unmarshal([]byte(str), &map_a)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("map反序列化后=%v\n", map_a)
}

func UnMarshalSlice() {

	var slice_a []map[string]interface{}
	//注意:反序列化 map,不需要 make,因为 make 操作被封装到 Unmarshal 函数
	str := "[{\"age\":12,\"name\":\"leighj\"},{\"age\":20,\"name\":\"tom\"}]"
	err := json.Unmarshal([]byte(str), &slice_a)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("slice反序列化后=%v\n", slice_a)
}

func reflectTest01(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)

	nu := rVal.Interface().(int)
	fmt.Println(nu)
}
