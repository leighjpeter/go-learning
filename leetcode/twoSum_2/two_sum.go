package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	temp := make(map[int]int)
	for i, v := range nums {
		temp[v] = i
		a := target - v
		if i2, ok := temp[a]; ok {
			return []int{i2, i}
		}
	}
	return []int{-1, -1}
}

func main() {
	target := 18
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, target))
}
