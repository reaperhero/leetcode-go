package _0034_计算x的n次幂函数

import (
	"fmt"
	"testing"
)

// 时间复杂度 O(log n),空间复杂度 O(1)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	tmp := myPow(x, n/2) // 除法是向下取整的 5 / 2 = 2
	if n%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x
}

func Test_myPow(t *testing.T) {
	fmt.Println(myPow(2, 5)) // 2的5次方 = 32
}
