package _6_partition_list

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	headDummy := &ListNode{Val: math.MinInt}
	headDummy.Next = head
	head = headDummy
	tailDummy := &ListNode{Val: math.MinInt}
	tail := tailDummy
	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			tail.Next = head.Next
			tail = tail.Next
			head.Next = head.Next.Next
		}
	}
	head.Next = tailDummy.Next
	tail.Next = nil
	return headDummy.Next
}
