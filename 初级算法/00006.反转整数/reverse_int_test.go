package _0006_反转整数

import (
	"fmt"
	"testing"
)

// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。注意:假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0
//Input: 120
//Output: 21

func reverse(x int) int {
	// int32 Range: -2147483648 through 2147483647.
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}

func Test_reverse(t *testing.T) {
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(120))
	fmt.Println(reverse(13466))
}
