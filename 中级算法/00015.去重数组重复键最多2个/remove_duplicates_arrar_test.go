package _0015_去重数组重复键最多2个

import (
	"fmt"
	"testing"
)

//给定一个有序数组 nums，对数组中的元素进行去重，使得原数组中的每个元素最多暴露 2 个。最后返回去重以后数组的长度值。

func removeDuplicates80(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		startFinder := -1
		for nums[finder] == nums[last] {
			if startFinder == -1 || startFinder > finder {
				startFinder = finder
			}
			if finder == len(nums)-1 {
				break
			}
			finder++
		}
		if finder-startFinder >= 2 && nums[finder-1] == nums[last] && nums[finder] != nums[last] {
			nums[last+1] = nums[finder-1]
			nums[last+2] = nums[finder]
			last += 2
		} else {
			nums[last+1] = nums[finder]
			last++
		}
		if finder == len(nums)-1 {
			if nums[finder] != nums[last-1] {
				nums[last] = nums[finder]
			}
			return last + 1
		}
	}
	return last + 1
}

func Test_removeDuplicates80(t *testing.T) {
	input := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates80(input)) // 7
}
