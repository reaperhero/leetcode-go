package _0042_旋转升序排序的数组查询值

import (
	"fmt"
	"testing"
)

func search(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	low, high := 0, len(nums)-1 // 0,7
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return []int{mid}
		} else if nums[mid] > nums[low] { // 在数值大的一部分区间里
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // 在数值小的一部分区间里
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return nil
}

func Test_search(t *testing.T) {
	input := []int{2, 5, 6, 0, 0, 1, 2}
	fmt.Println(search(input, 6)) // [2]
	fmt.Println(search(input, 7)) // []
}
