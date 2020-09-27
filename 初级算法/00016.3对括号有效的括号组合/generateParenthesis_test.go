package _0016_3对括号有效的括号组合

import (
	"fmt"
	"testing"
)

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	findGenerateParenthesis(n, n, "", &res)
	return res
}

func findGenerateParenthesis(lindex, rindex int, str string, res *[]string) {
	if lindex == 0 && rindex == 0 {
		*res = append(*res, str)
		return
	}
	if lindex > 0 {
		findGenerateParenthesis(lindex-1, rindex, str+"(", res)
	}
	if rindex > 0 && lindex < rindex {
		findGenerateParenthesis(lindex, rindex-1, str+")", res)
	}
}

func Test_generateParenthesis(t *testing.T) {
	fmt.Println(generateParenthesis(3)) // [((())) (()()) (())() ()(()) ()()()]
}
