package _0021_子串在母串中出现的下标

import (
	"fmt"
	"strings"
	"testing"
)

// 解法一
func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}

// 解法二
func strStr1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func Test_strStr(t *testing.T) {
	inout := "abacbabc"
	fmt.Println(strStr(inout, "abc"))
}
