package main

import "fmt"

/*
基本思想：参考
        1、归并排序是建立在归并操作上的一种有效的排序算法。
		该算法是采用分治法的一个非常典型的应用。
        ２、首先考虑下如何将2个有序数列合并。这个非常简单，
		只要从比较2个数列的第一个数，谁小就先取谁，取了后就在对应数列中删除这个数。
		然后再进行比较，如果有数列为空，那直接将另一个数列的数据依次取出即可。
---------------------------------------------------------------------------------
*/
func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge_(mergeSort(left), mergeSort(right))
}

func merge_(left []int, right []int) []int {
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

// 自顶向下归并排序，排序范围在 [begin,end) 的数组
func MergeSort(array []int, begin int, end int) {
	// 元素数量大于1时才进入递归
	if end-begin > 1 {
		// 将数组一分为二，分为 array[begin,mid) 和 array[mid,high)
		mid := begin + (end-begin+1)/2
		// 先将左边排序好
		MergeSort(array, begin, mid)
		// 再将右边排序好
		MergeSort(array, mid, end)
		// 两个有序数组进行合并
		merge(array, begin, mid, end)
	}
}
// 归并操作
func merge(array []int, begin int, mid int, end int) {
	// 申请额外的空间来合并两个有序数组，这两个数组是 array[begin,mid),array[mid,end)
	leftSize := mid - begin         // 左边数组的长度
	rightSize := end - mid          // 右边数组的长度
	newSize := leftSize + rightSize // 辅助数组的长度
	result := make([]int, 0, newSize)
	l, r := 0, 0
	for l < leftSize && r < rightSize {
		lValue := array[begin+l] // 左边数组的元素
		rValue := array[mid+r]   // 右边数组的元素
		// 小的元素先放进辅助数组里
		if lValue < rValue {
			result = append(result, lValue)
			l++
		} else {
			result = append(result, rValue)
			r++
		}
	}
	// 将剩下的元素追加到辅助数组后面
	result = append(result, array[begin+l:mid]...)
	result = append(result, array[mid+r:end]...)
	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < newSize; i++ {
		array[begin+i] = result[i]
	}
	return
}
func main() {
	list := []int{5}
	MergeSort(list, 0, len(list))
	fmt.Println(list)
	list1 := []int{5, 9}
	MergeSort(list1, 0, len(list1))
	fmt.Println(list1)
	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	MergeSort(list2, 0, len(list2))
	fmt.Println(list2)
}