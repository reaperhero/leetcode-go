package _0019_每k个元素翻转的方式翻转链表

import (
	"fmt"
	structures "leetcode-go/数据结构"
	"testing"
)

type ListNode = structures.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(first *ListNode, last *ListNode) *ListNode {
	prev := last
	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}
	return prev
}

func Test_reverseKGroup(t *testing.T) {
	input := structures.Ints2List([]int{3, 2, 1, 4, 5, 1, 34, 4, 6, 9, 23})
	result := structures.List2Ints(reverseKGroup(input, 3))
	fmt.Println(result) // [1 2 3 1 5 4 6 4 34 9 23]
}
