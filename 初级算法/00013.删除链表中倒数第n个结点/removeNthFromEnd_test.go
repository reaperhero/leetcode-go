package _0013_删除链表中倒数第n个结点

import (
	"fmt"
	"leetcode-go/数据结构"
	"testing"
)

type ListNode = structures.ListNode

// 解法一
// 先循环一次拿到链表的总长度，然后循环到要删除的结点的前一个结点开始删除操作。需要注意的一个特例是，有可能要删除头结点，要单独处理。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var fast, slow *ListNode
	fast = head
	slow = head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		head = head.Next
		return head
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

// 解法二
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n <= 0 {
		return head
	}
	current := head
	len := 0
	for current != nil {
		len++
		current = current.Next
	}
	if n > len {
		return head
	}
	if n == len {
		current := head
		head = head.Next
		current.Next = nil
		return head
	}
	current = head
	i := 0
	for current != nil {
		if i == len-n-1 {
			deleteNode := current.Next
			current.Next = current.Next.Next
			deleteNode.Next = nil
			break
		}
		i++
		current = current.Next
	}
	return head
}

func Test_removeNthFromEnd(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := removeNthFromEnd(structures.Ints2List(input), 3)
	fmt.Println(structures.List2Ints(result)) // [1 2 4 5]
}
