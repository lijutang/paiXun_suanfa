package main


/*
八、计数排序
	计数排序的核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。
	算法的步骤如下：
	根据待排序集合中最大元素和最小元素的差值范围，申请额外空间；
	遍历待排序集合，将每一个元素出现的次数记录到元素值对应的额外空间内；
	对额外空间内数据进行计算，得出每一个元素的正确位置；
	将待排序集合每一个元素移动到计算得出的正确位置上。
-------------------------------------------------------------------------------
*/
func countingSort(list []int, maxValue int) []int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen) // 初始为0的数组
	sortedIndex := 0
	length := len(list)
	for i := 0; i < length; i++ {
		bucket[list[i]] += 1
	}
	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			list[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}
	return list
}