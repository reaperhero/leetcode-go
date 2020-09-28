package _0039_升序二维矩阵查询值

import (
	"fmt"
	"testing"
)

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值,返回其索引值。该矩阵具有如下特性：
//
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数

func searchMatrix(matrix [][]int, target int) []int {
	if len(matrix) == 0 {
		return nil
	}
	m, low, high := len(matrix[0]), 0, len(matrix[0])*len(matrix)-1 // 4,0,11
	for low <= high {
		mid := low + (high-low)>>1          //5
		if matrix[mid/m][mid%m] == target { // 5/4=1 5%4=1
			return []int{mid / m, mid % m}
		} else if matrix[mid/m][mid%m] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return nil
}

func Test_(t *testing.T) {
	input := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	fmt.Println(searchMatrix(input, 16)) // [1 2]
}
