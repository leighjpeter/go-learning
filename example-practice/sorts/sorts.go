package main

import (
	"fmt"
)

// 冒泡排序-稳定
// 思路分析：在要排序的一组数中，对当前还未排好的序列，从前往后对相邻的两个数依次进行比较和调整，让较大的数往下沉，较小的往上冒。即，每当两相邻的数比较后发现它们的排序与排序要求相反时，就将它们互换。
// 时间复杂度O(n^2)
func BubbleSort() {
	a := [6]int{5, 2, 5, 7, 4, 8}
	fmt.Printf("冒泡排序前:%v\n", a)
	cnt := len(a)
	for i := 1; i < cnt; i++ {
		for j := 0; j < cnt-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Printf("冒泡排序后:%v\n", a)
}

// 选择排序-不稳定
// 思路分析：在要排序的一组数中，选出最小的一个数与第一个位置的数交换。然后在剩下的数当中再找最小的与第二个位置的数交换，如此循环到倒数第二个数和最后一个数比较为止。
// 时间复杂度O(n^2)
func SelectSort() {
	a := []int{5, 2, 5, 7, 4, 8}
	fmt.Printf("选择排序前:%v\n", a)
	cnt := len(a)
	for i := 0; i < cnt-1; i++ {
		min := a[i]
		minIndex := i
		for j := i + 1; j < cnt; j++ {
			if min >= a[j] {
				minIndex = j
				min = a[j]
			}
		}
		if minIndex != i {
			a[i], a[minIndex] = a[minIndex], a[i]
		}
	}
	fmt.Printf("选择排序后:%v\n", a)
}

// 插入排序-稳定
// 思路分析：在要排序的一组数中，假设前面的数已经是排好顺序的，现在要把第n个数插到前面的有序数中，使得这n个数也是排好顺序的。如此反复循环，直到全部排好顺序。
// 时间复杂度O(n^2)
func InsertSort() {
	a := []int{5, 2, 5, 7, 4, 8}
	fmt.Printf("插入排序前:%v\n", a)
	cnt := len(a)
	for i := 1; i < cnt; i++ {
		tmp := a[i]
		j := i - 1
		for ; j >= 0 && tmp < a[j]; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = tmp
	}
	fmt.Printf("插入排序后:%v\n", a)
}

// 快速排序-不稳定
// 思路分析：选择一个基准元素，通常选择第一个元素或者最后一个元素。通过一趟扫描，将待排序列分成两部分，一部分比基准元素小，一部分大于等于基准元素。此时基准元素在其排好序后的正确位置，然后再用同样的方法递归地排序划分的两部分
// 时间复杂度O(nlgn)
func QuickSort(a []int, leftIndex, rightIndex int) {
	if leftIndex >= rightIndex {
		return
	}
	i, j := leftIndex, rightIndex
	base := a[i]
	for i != j {
		for a[j] >= base && j > i {
			j--
		}
		if i <= j {
			a[i] = a[j]
		}
		for a[i] < base && i < j {
			i++
		}
		if i < j {
			a[j] = a[i]
		}
	}
	a[i] = base
	QuickSort(a, leftIndex, i-1)
	QuickSort(a, i+1, rightIndex)

}

// 归并排序-稳定
// 分析思路: 不断的分隔，然后两两比较排序，然后合并
// 时间复杂度O(nlgn)
func MergeSort(mergeSlice []int) []int {
	cnt := len(mergeSlice)
	if cnt < 2 {
		return mergeSlice
	}
	middle := cnt / 2

	left := mergeSlice[0:middle]
	right := mergeSlice[middle:]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

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

// 有序数组的二分查找
func BinaryFind(arr *[6]int, leftIndex, rightIndex, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("not find")
	}

	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Println("find!")
	}
}

// 汉诺塔游戏
// hannuota(3, "A", "B", "C")
func hannuota(n int, A, B, C string) {
	if n < 1 || n > 20 {
		return
	}
	if n == 1 {
		fmt.Printf("盘子%d从%s柱子移动到%s柱子\n", n, A, C)
	} else {
		hannuota(n-1, A, C, B)
		fmt.Printf("盘子%d从%s柱子移动到%s柱子\n", n, A, C)
		hannuota(n-1, B, A, C)
	}
}

var UnSortSlice = []int{6, 1, 2, 4, 9, 3, 7, 5, 10, 8, 0}
var heapSlice = []int{20, 30, 90, 40, 70, 110, 60, 10, 100, 50, 80}

func main() {
	// BubbleSort()
	// SelectSort()
	// InsertSort()

	// fmt.Printf("递归排序前:%v\n", UnSortSlice)
	// mergeSliceSort := MergeSort(UnSortSlice)
	// fmt.Printf("递归排序后:%v\n", mergeSliceSort)

	// fmt.Printf("快速排序前:%v\n", UnSortSlice)
	// QuickSort(UnSortSlice, 0, len(UnSortSlice)-1)
	// fmt.Printf("快速排序后:%v\n", UnSortSlice)

	// fmt.Printf("堆排序前:%v\n", heapSlice)
	// heapSliceSort := MergeSort(heapSlice)
	// fmt.Printf("堆排序后:%v\n", heapSliceSort)

	var val int
	// println(&val)
	fmt.Println(&val)
	f(10000)
	// println(&val)
	fmt.Println(&val)

	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s) // [0,0,0,0,0,1,2,3]
	ss := make([]int, 0)
	ss = append(ss, 1, 2, 3, 4)
	fmt.Println(ss) // [1,2,3,4]

	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])

	x := []int{100, 200, 300, 400, 500, 600, 700}
	twohundred := &x[1]
	x = append(x, 800)
	for i := range x {
		x[i]++
	}
	fmt.Println(*twohundred)

	intmap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}

	v, err := GetValue(intmap, 3)
	fmt.Println(v, err)
}

func f(i int) {
	if i--; i == 0 {
		return
	}
	f(i)
}
func hello(num ...int) {
	num[0] = 18
}

func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return "", false
}
