package main

import "fmt"

/*
1. 算法步骤
	一.选择一个增量序列 t1，t2，……，tk，其中 ti > tj, tk = 1；
	二.按增量序列个数 k，对序列进行 k 趟排序；
	三.每趟排序，根据对应的增量 ti，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。
	四.仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
*/

func shellSort(arr []int)  []int{
	length := len(arr)
	gap :=1
	for gap<gap/3 {
		gap = gap*3 +1
	}
	for gap>0 {
		for i := gap;i<length;i++ {
			temp := arr[i]
			j := i - gap
			for j >=0 && arr[j] >temp {
				arr[j+gap] = arr[j]
				j -=gap
			}
			arr[j+gap] = temp
		}
		gap = gap/3
	}
	return arr
}

func main() {
	array := []int{10,60,30,100,50}
	arr := shellSort(array)
	fmt.Println(arr)
}