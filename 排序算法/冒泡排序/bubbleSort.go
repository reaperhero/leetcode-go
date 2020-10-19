package main

import (
	"fmt"
)

//比较相邻的元素。如果第一个比第二个大，就交换它们两个；
//对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
//针对所有的元素重复以上的步骤，除了最后一个；
//重复步骤1~3，直到排序完成。

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func main()  {
	intslice := []int{2,4,6,2,8,1,5,6}
	bubbleSort(intslice)
	fmt.Println(intslice)
}


