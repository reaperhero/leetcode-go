package _0005_列出第n个顺序排列

import (
	"fmt"
	"strconv"
	"testing"
)

func getPermutation(n int, k int) string {
	if k == 0 {
		return ""
	}
	used, p, res := make([]bool, n), []int{}, ""
	findPermutation(n, 0, &k, p, &res, &used)
	return res
}

func findPermutation(n, index int, k *int, p []int, res *string, used *[]bool) {
	fmt.Printf("n = %v index = %v k = %v p = %v res = %v user = %v\n", n, index, *k, p, *res, *used)
	if index == n {
		*k--
		if *k == 0 {
			for _, v := range p {
				*res += strconv.Itoa(v + 1)
			}
		}
		return
	}
	for i := 0; i < n; i++ {
		if !(*used)[i] {
			(*used)[i] = true
			p = append(p, i)
			findPermutation(n, index+1, k, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	return
}


func Test_getPermutation(t *testing.T) {
	fmt.Println(getPermutation(4,5))

	fmt.Println(getPermutation(3,3)) // 213
	// "123"
	// "132"
	// "213"
	// "231"
	// "312"
	// "321"
}