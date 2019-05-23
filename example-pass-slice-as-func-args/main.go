package main

import (
	"fmt"
	"strings"
)

func modify(foo []string) {
	foo[0] = "b"
	fmt.Println("modify foo:", foo)
}

func addValue(foo []string) []string {
	foo = append(foo, "c")
	fmt.Println("modify foo:", foo)
	return foo
}

// slice 按址传递
func main() {
	foo := []string{"a", "b"}
	fmt.Println("before foo:", foo)
	// modify(foo)
	// foo = addValue(foo)
	fmt.Println("after foo:", foo)

	bar := foo[:1]
	fmt.Println("bar:", bar)

	s1 := append(bar, "c")
	fmt.Println("foo:", foo)
	fmt.Println("s1:", s1)

	s2 := append(bar, "d")
	fmt.Println("foo:", foo)
	fmt.Println("s2:", s2)

	// 如果超过原始的长度，则按值传递
	s3 := append(bar, "e", "f")
	fmt.Println("foo:", foo)
	fmt.Println("s3:", s3)
	// 	foo: [a d]
	// s3: [a e f]

	f := makeSuffix(".jpg")
	fmt.Println(f("winter"))
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
