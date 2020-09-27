package _0031_找到缺失的第一个正整数

import (
	"fmt"
	"testing"
)

// Input: [1,2,0]
// Output: 3

//Input: [3,4,-1,1]
//Output: 2

func firstMissingPositive(nums []int) int {
	numMap := make(map[int]int, len(nums))
	for _, v := range nums {
		numMap[v] = v
	}
	for index := 1; index < len(nums)+1; index++ {
		if _, ok := numMap[index]; !ok {
			return index
		}
	}
	return len(nums) + 1
}

func Test_firstMissingPositive(t *testing.T) {
	input1 := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(input1)) // 3

	input2 := []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(input2)) // 2
}
