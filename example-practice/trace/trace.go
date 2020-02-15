package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// (最大)堆排序-不稳定
// 思路分析：利用堆的特点 索引为i的节点的左子叶的索引是 2*i+1，右子叶的索引是 2*i+2，索引为i的父节点的索引是floor((i-1)/2)
func HeapSort(heapSlice []int) []int {
	cnt := len(heapSlice)
	// 从(n/2-1) --> 0逐次遍历。遍历之后，得到的数组实际上是一个(最大)二叉堆。
	for i := cnt/2 - 1; i >= 0; i-- {
		heapAdjust(heapSlice, i, cnt-1)
	}
	// 从最后一个元素开始对序列进行调整，不断的缩小调整的范围直到第一个元素
	for i := cnt - 1; i > 0; i-- {
		// 交换a[0]和a[i]。交换后，a[i]是a[0...i]中最大的。
		heapSlice[0], heapSlice[i] = heapSlice[i], heapSlice[0]
		// 调整a[0...i-1]，使得a[0...i-1]仍然是一个最大堆。
		// 即，保证a[i-1]是a[0...i-1]中的最大值。
		heapAdjust(heapSlice, 0, i-1)
	}
	return heapSlice
}

// 堆筛选，除了start之外，start~end均满足大顶堆的定义。
// 调整之后start~end称为一个大顶堆。
// @param arr 待调整数组
// @param start 起始指针
// @param end 结束指针
func heapAdjust(a []int, start, end int) {
	temp := a[start]
	// i是左叶子 i+1是右叶子
	for i := 2*start + 1; i <= end; i *= 2 {
		if i < end && a[i] < a[i+1] { // 左右叶子选择较大的
			i++
		}
		if temp >= a[i] {
			break
		} else { // 交换
			temp, a[i] = a[i], temp
		}
	}
}

var heapSlice = []int{20, 30, 90, 40, 70, 110, 60, 10, 100, 50, 80}

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	fmt.Printf("堆排序前:%v\n", heapSlice)
	heapSliceSort := HeapSort(heapSlice)
	fmt.Printf("堆排序后:%v\n", heapSliceSort)
}
