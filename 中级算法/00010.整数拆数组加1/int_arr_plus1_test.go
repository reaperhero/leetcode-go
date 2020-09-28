package _0010_整数拆数组加1

import (
	"fmt"
	"testing"
)

//Input: [1,2,3]
//Output: [1,2,4] 123+1=124
//Explanation: The array represents the integer 123.

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+carry > 9 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i] += carry
			carry = 0
		}
	}
	if digits[0] == 0 && carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func Test_plusOne(t *testing.T) {
	fmt.Println(plusOne([]int{1, 2, 3}))  // [1 2 4]
	fmt.Println(plusOne([]int{3, 2, 9}))  // [3 3 0]
	fmt.Println(plusOne([]int{3, -2, 9})) // [3 -1 0]
}
