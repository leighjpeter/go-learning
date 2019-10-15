package main

import (
	"fmt"
	"reflect"
)

const const_a int = 1
const const_b = 'A'
const const_c, const_d, const_e = 1, 2, 3

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

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

func main() {
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

	//================Reflect==================//
	/*
		//1. 先定义一个 int //
		var num int = 100
		reflectTest01(num)
	*/

}

func reflectTest01(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)

	nu := rVal.Interface().(int)
	fmt.Println(nu)
}
