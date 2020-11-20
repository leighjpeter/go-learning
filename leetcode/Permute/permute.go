package main

import "fmt"

/**
* 全排列
* 递归的公式 就是从数组中拿出一个，追加到剩余数组的全排列中
 */
func permute(arr []int) [][]int {
	if len(arr) == 1 {
		return [][]int{arr}
	}

	res := [][]int{}
	for i, num := range arr {
		tmp := make([]int, len(arr)-1) // 拿出一个元素之后的数组
		copy(tmp[0:], arr[0:i])
		copy(tmp[i:], arr[i+1:])
		sub := permute(tmp)
		for _, a := range sub {
			res = append(res, append(a, num))
		}
	}
	return res
}

func main() {
	arr := []int{1, 2, 3}
	for _, v := range permute(arr) {
		fmt.Println(v)
	}
}
