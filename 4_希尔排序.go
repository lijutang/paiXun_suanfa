package main

import "fmt"

/*
	基本思想：
        在要排序的一组数中，根据某一增量分为若干子序列，并对子序列分别进行插入排序。
        然后逐渐将增量减小,并重复上述过程。直至增量为1,此时数据序列基本有序,最后进行插入排序。
1. 算法步骤
	一.选择一个增量序列 t1，t2，……，tk，其中 ti > tj, tk = 1；
	二.按增量序列个数 k，对序列进行 k 趟排序；
	三.每趟排序，根据对应的增量 ti，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。
	四.仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
------------------------------------------------------------------------------------
*/

func shellSort(list []int)  []int{
	length := len(list)
	gap :=1
	for gap<gap/3 {
		gap = gap*3 +1
	}
	for gap>0 {
		for i := gap;i<length;i++ {
			temp := list[i]
			j := i - gap
			for j >=0 && list[j] >temp {
				list[j+gap] = list[j]
				j -=gap
			}
			list[j+gap] = temp
		}
		gap = gap/3
	}
	return list
}
//-------------------------------------------------------------------------------------
// 增量序列折半的希尔排序
func shellSort1(list []int) []int{
	// 数组长度
	n := len(list)
	// 每次减半，直到步长为 1
	for step := n / 2; step >= 1; step /= 2 {
		// 开始插入排序，每一轮的步长为 step
		for i := step; i < n; i += step {
			for j := i - step; j >= 0; j -= step {
				// 满足插入那么交换元素
				if list[j+step] < list[j] {
					list[j], list[j+step] = list[j+step], list[j]
					continue
				}
				break
			}
		}
	}
	return list
}
func main() {
	array := []int{10,60,30,100,50}
	//arr := shellSort(array)
	arr := shellSort1(array)
	fmt.Println(arr)
}