package _0041_两个整数之间k个数字整合

import (
	"fmt"
	"testing"
)

//
//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//Input: n = 4, k = 2
//Output:
//[
//[2,4],
//[3,4],
//[2,3],
//[1,2],
//[1,3],
//[1,4],
//]

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	generateCombinations(n, k, 1, c, &res) // 4,2,1,[]int,[][]int
	return res
}

func generateCombinations(n, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	// i will at most be n - (k - c.size()) + 1
	for i := start; i <= n-(k-len(c))+1; i++ { // 1,i<4-(2-0)+1,i++
		c = append(c, i)
		generateCombinations(n, k, i+1, c, res)
		c = c[:len(c)-1]
	}
	return
}

func Test_combine(t *testing.T) {
	input1, input2 := 4, 2
	fmt.Println(combine(input1, input2)) // [[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]
}
