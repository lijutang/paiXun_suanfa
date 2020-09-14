package main

import "fmt"

/*
 基本思想：
	在长度为N的无序数组中，第一次遍历n-1个数，找到最小的数值与第一个元素交换；
	第二次遍历n-2个数，找到最小的数值与第二个元素交换；
	。。。
	第n-1次遍历，找到最小的数值与第n-1个元素交换，排序完成。
1.算法步骤
	一.首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
	二.再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
	三.重复第二步，直到所有元素均排序完毕。
----------------------------------------------------------------------------------
*/

func selectionSort(list []int) []int {
	lengh := len(list)
	for i :=0;i<lengh-1;i++ {
		min := i
		for j :=i+1;j<lengh;j++ {
			if list[min]>list[j] {
				min = j
			}
		}
		list[i],list[min] = list[min],list[i]
	}
	return list
}

//--------------------------------------------------------------------------------
func selectionSort1(list []int) []int {
	n := len(list)
	//进行N-1轮迭代
	for i:=0;i<n-1;i++ {
		//每次从第ｉ位开始，找到最小的元素
		min := list[i] //最小数
		minIndex := i //最小的下标
		for j:=i+1;j<n;j++ {
			if  list[j]<min {
				// 如果找到的数比上次的还小，那么最小的数变为它
				min = list[j]
				minIndex = j
			}
		}
		// 这一轮找到的最小数的下标不等于最开始的下标，交换元素
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}

	}
	return list
}
//---------------------------------------------------------------------------------
func selectionSort2(list []int)   []int{
	n := len(list)
	// 只需循环一半
	for i := 0; i < n/2; i++ {
		minIndex := i // 最小值下标
		maxIndex := i // 最大值下标
		// 在这一轮迭代中要找到最大值和最小值的下标
		for j := i + 1; j < n-i; j++ {
			// 找到最大值下标
			if list[j] > list[maxIndex] {
				maxIndex = j // 这一轮这个是大的，直接 continue
				continue
			}
			// 找到最小值下标
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		if maxIndex == i && minIndex != n-i-1 {
			// 如果最大值是开头的元素，而最小值不是最尾的元素
			// 先将最大值和最尾的元素交换
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
			// 然后最小的元素放在最开头
			list[i], list[minIndex] = list[minIndex], list[i]
		} else if maxIndex == i && minIndex == n-i-1 {
			// 如果最大值在开头，最小值在结尾，直接交换
			list[minIndex], list[maxIndex] = list[maxIndex], list[minIndex]
		} else {
			// 否则先将最小值放在开头，再将最大值放在结尾
			list[i], list[minIndex] = list[minIndex], list[i]
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
		}
	}
	return list
}
func main() {
	array := []int{10,60,30,100,50}
	//arr := selectionSort(array)
	//arr := selectionSort1(array)
	arr := selectionSort2(array)
	fmt.Println(arr)
}