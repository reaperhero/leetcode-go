package _0017_比x大的数字都往后面排列

import (
	"fmt"
	structures "leetcode-go/数据结构"
	"testing"
)

//给定一个数 x，比 x 大或等于的数字都要排列在比 x 小的数字后面，并且相对位置不能发生变化。由于相对位置不能发生变化，所以不能用类似冒泡排序的思想。
//Input: head = 1->4->3->2->5->2, x = 3
//Output: 1->2->2->4->3->5

// ListNode define
type ListNode = structures.ListNode

// 解法一 单链表
// [1, 4, 3, 2, 5, 2]
func partition(head *ListNode, x int) *ListNode {
	beforeHead := &ListNode{Val: 0, Next: nil}
	before := beforeHead
	afterHead := &ListNode{Val: 0, Next: nil}
	after := afterHead

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterHead.Next
	return beforeHead.Next
}

// DoublyListNode define
type DoublyListNode struct {
	Val  int
	Prev *DoublyListNode
	Next *DoublyListNode
}

// 解法二 双链表
func partition1(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	DLNHead := genDoublyListNode(head)
	cur := DLNHead
	for cur != nil {
		if cur.Val < x {
			tmp := &DoublyListNode{Val: cur.Val, Prev: nil, Next: nil}
			compareNode := cur
			for compareNode.Prev != nil {
				if compareNode.Val >= x && compareNode.Prev.Val < x {
					break
				}
				compareNode = compareNode.Prev
			}
			if compareNode == DLNHead {
				if compareNode.Val < x {
					cur = cur.Next
					continue
				} else {
					tmp.Next = DLNHead
					DLNHead.Prev = tmp
					DLNHead = tmp
				}
			} else {
				tmp.Next = compareNode
				tmp.Prev = compareNode.Prev
				compareNode.Prev.Next = tmp
				compareNode.Prev = tmp
			}
			deleteNode := cur
			if cur.Prev != nil {
				deleteNode.Prev.Next = deleteNode.Next
			}
			if cur.Next != nil {
				deleteNode.Next.Prev = deleteNode.Prev
			}
		}
		cur = cur.Next
	}
	return genListNode(DLNHead)
}

func genDoublyListNode(head *ListNode) *DoublyListNode {
	cur := head.Next
	DLNHead := &DoublyListNode{Val: head.Val, Prev: nil, Next: nil}
	curDLN := DLNHead
	for cur != nil {
		tmp := &DoublyListNode{Val: cur.Val, Prev: curDLN, Next: nil}
		curDLN.Next = tmp
		curDLN = tmp
		cur = cur.Next
	}
	return DLNHead
}

func genListNode(head *DoublyListNode) *ListNode {
	cur := head.Next
	LNHead := &ListNode{Val: head.Val, Next: nil}
	curLN := LNHead
	for cur != nil {
		tmp := &ListNode{Val: cur.Val, Next: nil}
		curLN.Next = tmp
		curLN = tmp
		cur = cur.Next
	}
	return LNHead
}

func Test_partition(t *testing.T) {
	input1, input2 := structures.Ints2List([]int{1, 4, 3, 2, 5, 2}), 3
	fmt.Println(structures.List2Ints(partition(input1, input2))) // [1 2 2 4 3 5]
}
