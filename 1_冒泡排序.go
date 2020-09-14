package main

import "fmt"

/*
*1. 算法步骤
	一.比较相邻的元素。如果第一个比第二个大，就交换他们两个。
	二.对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
	三.针对所有的元素重复以上的步骤，除了最后一个。
	四.持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
*
*/

func bubbleSort(arr []int) []int {
	length := len(arr)
	for i :=0;i<length;i++ {
		for j := 0;j<length-i-1;j++ {
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	return arr
}

func main() {
	array := []int{10,60,30,100,50}
	arr := bubbleSort(array)
	fmt.Println(arr)
}