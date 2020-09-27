package _0002_顺时针旋转90度二位数组

import (
	"fmt"
	"testing"
)

func rotate(matrix [][]int) {
	row := len(matrix)
	if row <= 0 {
		return
	}
	column := len(matrix[0])
	// rotate by diagonal 对角线变换
	for i := 0; i < row; i++ {
		for j := i + 1; j < column; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
	// rotate by vertical centerline 竖直轴对称翻转
	halfColumn := column / 2
	for i := 0; i < row; i++ {
		for j := 0; j < halfColumn; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][column-j-1]
			matrix[i][column-j-1] = tmp
		}
	}
}

func Test_rotate(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for k, _ := range input {
		fmt.Println(input[k])
	}
	// [1 2 3]
	// [4 5 6]
	// [7 8 9]
	rotate(input)
	for k, _ := range input {
		fmt.Println(input[k])
	}
	// [7 4 1]
	// [8 5 2]
	// [9 6 3]
}
