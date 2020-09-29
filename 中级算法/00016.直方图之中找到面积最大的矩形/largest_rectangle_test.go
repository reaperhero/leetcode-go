package _0016_直方图之中找到面积最大的矩形

import (
	"fmt"
	"testing"
)

//给出每个直方图的高度，要求在这些直方图之中找到面积最大的矩形，输出矩形的面积。

//Input: [2,1,5,6,2,3]
//Output: 10

func largestRectangleArea(heights []int) int {
	maxArea, stack, height := 0, []int{}, 0
	for i := 0; i <= len(heights); i++ {
		if i == len(heights) {
			height = 0
		} else {
			height = heights[i] // 2
		}
		if len(stack) == 0 || height >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			tmp := stack[len(stack)-1]
			fmt.Printf("1. tmp = %v stack = %v\n", tmp, stack)
			stack = stack[:len(stack)-1]
			length := 0
			if len(stack) == 0 {
				length = i
			} else {
				length = i - 1 - stack[len(stack)-1]
				fmt.Printf("2. length = %v stack = %v i = %v\n", length, stack, i)
			}
			maxArea = max(maxArea, heights[tmp]*length)
			fmt.Printf("3. maxArea = %v heights[tmp]*length = %v\n", maxArea, heights[tmp]*length)
			i--
		}
	}
	return maxArea
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_largestRectangleArea(t *testing.T) {
	input := []int{1, 2, 1, 3}
	fmt.Println(largestRectangleArea(input)) // 4
}
