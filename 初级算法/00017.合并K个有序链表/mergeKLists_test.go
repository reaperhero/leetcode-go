package _0017_合并K个有序链表

import (
	"fmt"
	structures "leetcode-go/数据结构"
	"testing"
)

type ListNode = structures.ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeTwoLists1(left, right)
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists1(l1, l2.Next)
	return l2
}

func Test_(t *testing.T) {
	input1 := structures.Ints2List([]int{1, 3, 8})
	input2 := structures.Ints2List([]int{2, 3, 4, 4, 5, 6})
	input3 := structures.Ints2List([]int{1, 9, 9, 9, 9, 9})
	reslut := mergeTwoLists1(mergeTwoLists1(input1, input2), input3)
	fmt.Println(structures.List2Ints(reslut)) // [1 1 2 3 3 4 4 5 6 8 9 9 9 9 9]
}
