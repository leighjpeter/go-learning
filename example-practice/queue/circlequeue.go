package main

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

// add queue
func (this *CircleQueue) Push(val int) (err error) {
	if this.isFull() {
		return errors.New("queue full")
	}

	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

// get
func (this *CircleQueue) Pop() (val int, err error) {
	if this.isEmpty() {
		return 0, errors.New("Empty")
	}

	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

func (this *CircleQueue) isFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

func (this *CircleQueue) isEmpty() bool {
	return this.tail == this.head
}

func (this *CircleQueue) ListQueue() {
	num := this.GetNum()
	if num == 0 {
		fmt.Println("Empty")
	}
	tempHead := this.head
	for i := 0; i < num; i++ {
		fmt.Printf("arr[%d]=%d\n", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
}

// get count
func (this *CircleQueue) GetNum() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {
	cq := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	cq.Push(1)
	cq.Push(2)
	cq.Push(3)
	err := cq.Push(4)
	if err != nil {
		fmt.Println(err)
	}
	cq.ListQueue()
	v, _ := cq.Pop()
	cq.Pop()
	cq.Pop()
	cq.Pop()
	fmt.Println(v)
	cq.ListQueue()
	err = cq.Push(5)
	if err != nil {
		fmt.Println(err)
	}
	err = cq.Push(6)
	if err != nil {
		fmt.Println(err)
	}
	err = cq.Push(7)
	if err != nil {
		fmt.Println(err)
	}
	cq.ListQueue()
}
