package main

import "fmt"

/*
1.算法步骤
	一.首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
	二.再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
	三.重复第二步，直到所有元素均排序完毕。
*/

func selectionSort(arr []int) []int {
	lengh := len(arr)
	for i :=0;i<lengh-1;i++ {
		min := i
		for j :=i+1;j<lengh;j++ {
			if arr[min]>arr[j] {
				min = j
			}
		}
		arr[i],arr[min] = arr[min],arr[i]
	}
	return arr
}



func main() {
	array := []int{10,60,30,100,50}
	arr := selectionSort(array)
	fmt.Println(arr)
}