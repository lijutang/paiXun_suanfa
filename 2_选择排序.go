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

func SelectSortMax(arr[] int ) int {
	length := len(arr) // 数组长度
	if length<=1 {
		return arr[0] //一个元素的数组直接返回
	} else {
		max := arr[0]
		for i := 1 ;i <length;i++ {
			if arr[i] >max { //任何以比我大的数，是最大值
				max = arr[i]
			}
		}
		return max
	}
}

func SelectSort(arr[] int ) []int {
	length := len(arr) // 数组长度
	if length<=1 {
		return arr //一个元素的数组直接返回
	} else {
		for i := 0 ;i<length-1 ;i++ {// 进行 N-1 轮迭代(只剩一个元素的时候不需要挑选)
			// 每次从第 i 位开始，找到最小的元素
			min :=i //标记索引
			for j:= i+1;j<length;j++ { // 每次选出一个极小值
				if arr[min]>arr[j] {
					min = j  //保存极小值的索引
				}
			}
			if i != min {
				arr[i],arr[min] = arr[min],arr[i] // 数据交换
			}
		}
		return arr
	}


}

func main() {
	array := []int{10,60,30,100,50}
	arr := SelectSort(array)

	fmt.Println(arr)
}