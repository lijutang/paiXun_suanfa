package main

import "fmt"

/*
	算法描述：首先建一个堆，然后调整堆，调整过程是将节点和子节点进行比较，
	将 其中最大的值变为父节点，递归调整调整次数lgn,最后将根节点和尾节点交换再n次 调整O(nlgn).
	算法步骤
	创建最大堆或者最小堆（我是最小堆）
	调整堆
	交换首尾节点(为了维持一个完全二叉树才要进行收尾交换)
------------------------------------------------------------------------------------
*/



func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(HeapSortMax(arr,len(arr)))
}
func HeapSortMax(arr []int, length int) []int {
	// length := len(arr)
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1 //二叉树深度，n 2*n+1 2*n+2
	for i := depth; i >= 0; i-- {
		topmax := i //假定最大的位置就在i的位置
		leftchild := 2*i + 1
		rightchild := 2*i + 2 // 左右孩子节点。
		if leftchild <= length-1 && arr[leftchild] > arr[topmax] { //防止越过界限，如果左边比我大，记录最大
			topmax = leftchild
		}
		if rightchild <= length-1 && arr[rightchild] > arr[topmax] { //防止越过界限，如果右边比我大，记录最大
			topmax = rightchild
		}
		if topmax != i {//确保i的值就是最大
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}
	return arr
}
func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastlen := length - i // 每次截取一段
		HeapSortMax(arr, lastlen)
		if i < length {
			arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
		}
	}
	return arr
}


