package _0011_12步爬楼梯

import (
	"fmt"
	"testing"
)

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？注意：给定 n 是一个正整数
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func Test_climbStairs(t *testing.T) {
	fmt.Println(climbStairs(1))  //1
	fmt.Println(climbStairs(2))  //2
	fmt.Println(climbStairs(3))  //3
	fmt.Println(climbStairs(10)) //89
}
