package leetcode

import (
	"fmt"
	"testing"
)

func removeDuplicates(source []int) []int {
	if len(source) < 1 {
		return source
	}
	var checkMap = make(map[int]int)
	for i := 0; i < len(source); i++ {
		checkValue := source[i]
		if _, ok := checkMap[checkValue]; ok {
			source = append(source[:i], source[i+1:]...)
			i--
			continue
		} else {
			checkMap[checkValue] = i
		}
	}
	return source
}

func Test_removeDuplicates(t *testing.T) {
	source := []int{1, 2, 1, 4, 5, 6, 7, 10, 1}
	fmt.Println(removeDuplicates(source)) // [1 2 4 5 6 7 10]
}
