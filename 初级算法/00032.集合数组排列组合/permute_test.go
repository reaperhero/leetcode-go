package _0032_集合数组排列组合

import (
	"fmt"
	"testing"
)

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
	generatePermutation(nums, 0, p, &res, &used)
	return res
}

func generatePermutation(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			p = append(p, nums[i])
			generatePermutation(nums, index+1, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	return
}

func Test_permute(t *testing.T) {
	input := []int{1, 2, 3}
	fmt.Println(permute(input)) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}
