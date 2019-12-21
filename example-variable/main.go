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

type person struct {
	nickname string
	sex      string
	birthday string
	weight   string
	sayHello string
}

type IFile interface {
	Read(buf []byte) (err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}

type IReader interface {
	Read(buf []byte) (n int, err error)
}

type IWriter interface {
	Write(buf []byte) (n int, err error)
}

type ICloser interface {
	Close() error
}

type File struct {
	// ...
}

func (f *File) Read(buf []byte) (n int, err error)                { return }
func (f *File) Write(buf []byte) (n int, err error)               { return }
func (f *File) Seek(off int64, whence int) (pos int64, err error) { return }
func (f *File) Close() error                                      { return nil }

func main() {
	//================SLICE=======================//

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

	//================Reflect==================//

	//1. 先定义一个 int //
	var num int = 100
	reflectTest01(num)

}

func reflectTest01(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)

	nu := rVal.Interface().(int)
	fmt.Println(nu)
}
