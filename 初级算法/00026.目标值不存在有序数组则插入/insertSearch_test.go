package _0026_目标值不存在有序数组则插入

import (
	"fmt"
	"testing"
)

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] >= target {
			high = mid - 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			low = mid + 1
		}
	}
	return 0
}

func Test_searchInsert(t *testing.T) {
	input := []int{1,3,5,6,7,8,9,10,12,16}
	fmt.Println(searchInsert(input,2))
}