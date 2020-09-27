package _0020_数组中所有等于val的元素删除

import (
	"fmt"
	"testing"
)

// 给定一个数组 nums 和一个数值 val，将数组中所有等于 val 的元素删除，并返回剩余的元素个数。
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}
	return j
}

func Test_removeElement(t *testing.T) {
	input := []int{0, 1, 0, 3, 0, 12}
	fmt.Println(removeElement(input, 0))
}
