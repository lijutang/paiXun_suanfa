package main

import "fmt"

/*
基本思想：
        在要排序的一组数中，假定前n-1个数已经排好序，
	现在将第n个数插到前面的有序数列中，使得这n个数也是排好顺序的。
	如此反复循环，直到全部排好顺序。
算法步骤
一般来说，插入排序都采用in-place在数组上实现。具体算法描述如下：
	从第一个元素开始，该元素可以认为已经被排序；
	取出下一个元素，在已经排序的元素序列中从后向前扫描；
	如果该元素（已排序）大于新元素，将该元素移到下一位置；
	重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
	将新元素插入到该位置后；
	重复步骤2~5。
------------------------------------------------------------------------------------
*/


func insertionSort(list []int) []int {
	for i := range list {
		preIndex:= i-1
		current := list[i]
		for preIndex >=0 && list[preIndex]>current{
			list[preIndex+1] = list[preIndex]
			preIndex -= 1
		}
		list[preIndex+1] = current
	}
	return list
}
//------------------------------------------------------------------------------------
func insertionSort1(list []int) []int{
	n := len(list)
	// 进行 N-1 轮迭代
	for i := 1; i <= n-1; i++ {
		deal := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置
		// 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理
		if deal < list[j] {
			// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j] // 某数后移，给待排序留空位
			}
			list[j+1] = deal // 结束了，待排序的数插入空位
		}
	}
	return list
}
//------------------------------------------------------------------------------------
func main() {
	array := []int{10,60,30,100,50}
	//arr := insertionSort(array)
	arr := insertionSort1(array)
	fmt.Println(arr)
}