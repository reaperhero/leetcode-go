package _0044_删除重复的节点

import (
	"fmt"
	structures "leetcode-go/数据结构"
	"testing"
)

//删除链表中重复的结点，以保障每个结点只出现一次。

//Input: 1->1->2->3->3
//Output: 1->2->3

type ListNode = structures.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func Test_deleteDuplicates(t *testing.T) {
	input := structures.Ints2List([]int{1, 1, 2, 2, 3, 3, 3})
	result := structures.List2Ints(deleteDuplicates(input))
	fmt.Println(result) // [1 2 3]
}
