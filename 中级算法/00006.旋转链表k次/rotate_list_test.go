package _0006_旋转链表k次

import (
	"fmt"
	structures "leetcode-go/数据结构"
	"testing"
)

// 这道题需要注意的点是，K 可能很大，K = 2000000000 ，如果是循环肯定会超时。应该找出 O(n) 的复杂度的算法才行。由于是循环旋转，最终状态其实是确定的，利用链表的长度取余可以得到链表的最终旋转结果。

//Input: 1->2->3->4->5->NULL, k = 2
//Output: 4->5->1->2->3->NULL
//Explanation:
//rotate 1 steps to the right: 5->1->2->3->4->NULL
//rotate 2 steps to the right: 4->5->1->2->3->NULL

type ListNode = structures.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	len := 0
	cur := newHead
	for cur.Next != nil {
		len++
		cur = cur.Next
	}
	if (k % len) == 0 {
		return head
	}
	cur.Next = head
	cur = newHead
	for i := len - k%len; i > 0; i-- {
		cur = cur.Next
	}
	res := &ListNode{Val: 0, Next: cur.Next}
	cur.Next = nil
	return res.Next
}

func Test_rotateRight(t *testing.T) {
	input1 := structures.Ints2List([]int{1, 2, 3, 4, 5, 6, 7})
	result := structures.List2Ints(rotateRight(input1, 2))
	fmt.Println(result) // [6 7 1 2 3 4 5]
}
