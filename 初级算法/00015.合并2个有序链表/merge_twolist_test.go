package _0015_合并2个有序链表

import (
	"fmt"
	structures "leetcode-go/数据结构"
	"testing"
)

// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4

type ListNode = structures.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func Test_mergeTwoLists(t *testing.T) {
	input1 := structures.Ints2List([]int{1, 2, 4})
	input2 := structures.Ints2List([]int{1, 3, 4})
	result := structures.List2Ints(mergeTwoLists(input1, input2))
	fmt.Println(result)
}
