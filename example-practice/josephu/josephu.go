package main

import (
	"fmt"
)

type Boy struct {
	No   int
	Next *Boy
}

func AddBoy(num int) *Boy {
	first, cur := &Boy{}, &Boy{} // 空节点
	if num < 1 {
		fmt.Println("num error")
		return first
	}

	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		if i == 1 {
			first = boy
			cur = boy
			cur.Next = first
		} else {
			cur.Next = boy
			cur = boy
			cur.Next = first
		}
	}
	return first
}

func ShowBoy(first *Boy) {
	if first.Next == nil {
		fmt.Println("no boy")
	}
	cur := first
	for {
		fmt.Printf("小孩编号=%d,next=%v\n", cur.No, cur.Next)
		if cur.Next == first {
			break
		}
		cur = cur.Next
	}
}

func PlayGame(first *Boy, startNo int, countNum int) {
	tail := first
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}
	for i := 1; i < startNo; i++ {
		first = first.Next
		tail = tail.Next
	}
	for {
		for i := 1; i < countNum; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("编号%d的小孩出列\n", first.No)
		first = first.Next
		tail.Next = first
		// fmt.Println(tail, first)
		if tail == first {
			break
		}
	}
	fmt.Printf("编号%d的小孩出列\n", first.No)
}

func main() {
	first := AddBoy(5)
	ShowBoy(first)
	PlayGame(first, 2, 3)
}
