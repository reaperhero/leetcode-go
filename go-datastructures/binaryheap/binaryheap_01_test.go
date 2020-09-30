package binaryheap

import (
	"fmt"
	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
	"testing"
)

// 二叉堆（英语：binary heap）是一种特殊的堆，二叉堆是完全二叉树或者是近似完全二叉树
//
func Test_binaryheap(t *testing.T) {

	// 最小堆
	heap := binaryheap.NewWithIntComparator() // empty (min-heap)
	heap.Push(2)
	heap.Push(3)
	heap.Push(1)
	heap.Push(4)
	heap.Push(5)
	heap.Push(6)
	heap.Push(6)
	heap.Push(2)
	fmt.Println(heap.Values()) // [1 2 2 3 5 6 6 4]
	_, _ = heap.Peek()         // 1,true
	iteratorMin := heap.Iterator()
	for iteratorMin.Next() {
		fmt.Println(iteratorMin.Value()) // 从最小值开始打印
	}
	heap.Push(1) // 1
	heap.Clear() // empty
	heap.Empty() // true
	heap.Size()  // 0

	// 最大堆
	inverseIntComparator := func(a, b interface{}) int {
		return -utils.IntComparator(a, b)
	}
	heap = binaryheap.NewWith(inverseIntComparator) // empty (min-heap)
	heap.Push(2)                                    // 2
	heap.Push(3)                                    // 3, 2
	heap.Push(1)                                    // 3, 2, 1
	heap.Push(4)
	heap.Push(5)
	heap.Push(6)
	heap.Push(6)
	heap.Push(2)
	fmt.Println(heap.Values()) // [6 4 6 2 3 1 5 2]
	iteratorMax := heap.Iterator()
	for iteratorMax.Next() {
		fmt.Println(iteratorMax.Value()) // 从最大值开始打印
	}
}
