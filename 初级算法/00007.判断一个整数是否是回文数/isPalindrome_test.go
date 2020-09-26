package _0007_判断一个整数是否是回文数

import (
	"fmt"
	"strconv"
	"testing"
)

// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func Test_isPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(120))
	fmt.Println(isPalindrome(222))
}