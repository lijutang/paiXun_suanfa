package main

import "fmt"

/*

	基本思想：两个数比较大小，较大的数下沉，较小的数冒起来。
    过程：比较相邻的两个数据，如果第二个数小，就交换位置。
	  从后向前两两比较，一直到比较最前两个数据。最终最小数被交换到起始的位置，这样第一个最小数的位置就排好了。
	  继续重复上述过程，依次将第2.3...n-1个最小数排好位置。
-----------------------------------------------------------------------------------
*
*/

func BubbleSort(list [5]int) [5]int {
	length := len(list)
	for i := 0 ;i< length;i++ { //最外层表示一共循环的次数
		for j := 0; j < length-1-i; j++ { // //内层表示逐层比较的次数递减
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

/*优化：
	针对问题：
	数据的顺序排好之后，冒泡算法仍然会继续进行下一轮的比较，直到arr.length-1次，后面的比较没有意义的。
	方案：
	设置标志位flag，如果发生了交换flag设置为true；如果没有交换就设置为false。
	这样当一轮比较结束后如果flag仍为false，即：这一轮没有发生交换，说明数据的顺序已经排好，没有必要继续进行下去。
---------------------------------------------------------------------------------------
*/

func BubbleSort1(list [5]int) [5]int {
	n :=len(list)
	//在一轮中有没有交换过
	didSwap := false
	for i :=n-1;i>0;i-- {
		// 每次从第一位开始比较，比较到第 i 位就不比较了，因为前一轮该位已经有序了
		for j :=0;j<i;j++ {
			// 如果前面的数比后面的大，那么交换
			if list[j]>list[j+1] {
				list[j],list[j+1] = list[j+1],list[j]
				didSwap = true
			}
		}
		// 如果在一轮中没有交换过，那么已经排好序了，直接返回
		if !didSwap {
			return list
		}
	}
	return list
}

func BubbleSort2(list [5]int) [5]int {
	tmp := 0
	for i := 0; i < len(list); i++ {
		isSorted := true
		for j := 0; j < len(list)-1; j++ {
			if list[j] > list[j+1] {
				tmp = list[j]
				list[j] = list[j+1]
				list[j+1] = tmp
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
	return list
}


func main() {
	array := [5]int{10,60,30,100,50}
	//arr := BubbleSort(array)
	//arr := BubbleSort1(array)
	arr := BubbleSort2(array)
	fmt.Println(arr)
}