package _0037_二进制字符串相加

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}

	res := make([]string, len(a)+1)
	i, j, k, c := len(a)-1, len(b)-1, len(a), 0
	for i >= 0 && j >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		bj, _ := strconv.Atoi(string(b[j]))
		res[k] = strconv.Itoa((ai + bj + c) % 2)
		c = (ai + bj + c) / 2
		i--
		j--
		k--
	}

	for i >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		res[k] = strconv.Itoa((ai + c) % 2)
		c = (ai + c) / 2
		i--
		k--
	}

	if c > 0 {
		res[k] = strconv.Itoa(c)
	}

	return strings.Join(res, "")
}

func Test_addBinary(t *testing.T) {
	str1 := "1010"
	str2 := "1011"
	fmt.Println(addBinary(str1, str2)) // 10101
}
