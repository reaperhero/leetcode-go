package _0013_数组所有可能的子集

import (
	"fmt"
	"sort"
	"testing"
)

//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。说明：解集不能包含重复的子集。

//Input: nums = [1,2,3]
//Output:
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]

// 解法一
func subsets(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	for k := 0; k <= len(nums); k++ { // 生成0-len(nums)长度之间的数组
		generateSubsets(nums, k, 0, c, &res)
	}
	return res
}

func generateSubsets(nums []int, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	// i will at most be n - (k - c.size()) + 1
	for i := start; i < len(nums)-(k-len(c))+1; i++ { // i < 3 - (1-0)+1
		c = append(c, nums[i])
		generateSubsets(nums, k, i+1, c, res)
		c = c[:len(c)-1]
	}
	return
}

// 解法二
func subsets1(nums []int) [][]int {
	res := make([][]int, 1)
	sort.Ints(nums)
	for i := range nums {
		for _, org := range res {
			clone := make([]int, len(org), len(org)+1)
			copy(clone, org)
			clone = append(clone, nums[i])
			res = append(res, clone)
		}
	}
	return res
}

// 解法三：位运算的方法
func subsets2(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := [][]int{}
	sum := 1 << uint(len(nums))
	for i := 0; i < sum; i++ {
		stack := []int{}
		tmp := i                              // i 从 000...000 到 111...111
		for j := len(nums) - 1; j >= 0; j-- { // 遍历 i 的每一位
			if tmp&1 == 1 {
				stack = append([]int{nums[j]}, stack...)
			}
			tmp >>= 1
		}
		res = append(res, stack)
	}
	return res
}

func Test_subsets(t *testing.T) {
	input := []int{1, 2, 3}
	fmt.Println(subsets(input)) // [[] [1] [2] [3] [1 2] [1 3] [2 3] [1 2 3]]
}
