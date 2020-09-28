package _0040_数组排序

import (
	"fmt"
	"testing"
)

//Input: [2,0,2,1,1,0]
//Output: [0,0,1,1,2,2]

// 只包含0，1，2数组排序
func sortArray(nums []int) {
	if len(nums) == 0 {
		return
	}

	r := 0
	w := 0
	b := 0 // label the end of different colors;
	for _, num := range nums {
		if num == 0 {
			nums[b] = 2
			b++
			nums[w] = 1
			w++
			nums[r] = 0
			r++
		} else if num == 1 {
			nums[b] = 2
			b++
			nums[w] = 1
			w++
		} else if num == 2 {
			b++
		}
	}
}


func Test_sortArray(t *testing.T) {
	input := []int{2,0,2,1,1,0}
	sortArray(input)
	fmt.Println(input) // [0 0 1 1 2 2]
}