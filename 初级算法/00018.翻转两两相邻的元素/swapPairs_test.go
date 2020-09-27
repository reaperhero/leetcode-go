package _0018_翻转两两相邻的元素

import (
	"fmt"
	structures "leetcode-go/数据结构"
	"testing"
)

type ListNode = structures.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	s := head.Next
	var behind *ListNode
	for head.Next != nil {
		headNext := head.Next
		if behind != nil && behind.Next != nil {
			behind.Next = headNext
		}
		var next *ListNode
		if head.Next.Next != nil {
			next = head.Next.Next
		}
		if head.Next.Next != nil {
			head.Next = next
		} else {
			head.Next = nil
		}
		headNext.Next = head
		behind = head
		if head.Next != nil {
			head = next
		}
	}
	return s
}

func Test_swapPairs(t *testing.T) {
	input := structures.Ints2List([]int{1, 2, 3, 4, 5})
	reslut := structures.List2Ints(swapPairs(input))
	fmt.Println(reslut) // [2 1 4 3 5]
}
