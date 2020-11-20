package main

import "fmt"

type NodeList struct {
	Val  int
	Next *NodeList
}

func fromTailToHead(n *NodeList) {
	if n == nil {
		return
	}
	fromTailToHead(n.Next)
	fmt.Printf("%d ->", n.Val)
}

/*
从尾到头打印节点
*/
func main() {

	n3 := &NodeList{3, nil}
	n2 := &NodeList{2, n3}
	n1 := &NodeList{1, n2}

	fromTailToHead(n1)
	fmt.Printf("\n \n")
}
