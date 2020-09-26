package leetcode

import (
	"fmt"
	"testing"
)

func two_sum_00002(source []int, target int) []int {
	resultMap := make(map[int]int)
	for k, v := range source {
		if idx, ok := resultMap[target-v]; ok {
			return []int{idx, k}
		}
		resultMap[v] = k
	}
	return nil
}

func Test_two_sum_00002(t *testing.T) {
	result := two_sum_00002([]int{1, 2, 3, 5, 7}, 4)
	fmt.Println(result)
}
