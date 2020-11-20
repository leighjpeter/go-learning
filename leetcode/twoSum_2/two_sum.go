package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	temp := make(map[int]int)
	for i, v := range nums {
		a := target - v
		if i2, ok := temp[a]; ok {
			return []int{i2, i}
		}
		temp[v] = i
	}
	return []int{-1, -1}
}

func main() {
	target := 6
	nums := []int{3, 2, 4}
	fmt.Println(twoSum(nums, target))
}
