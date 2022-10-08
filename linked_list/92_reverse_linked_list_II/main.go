package _2_reverse_linked_list_II

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Val: math.MinInt}
	dummy.Next = head
	head = dummy
	var prev *ListNode
	for i := 0; i < left; i++ {
		prev = head
		head = head.Next
	}
	mid := head
	var next *ListNode
	for j := left; j <= right; j++ {
		if head == nil {
			break
		}
		temp := head.Next
		head.Next = next
		next = head
		head = temp
	}
	prev.Next = next
	mid.Next = head
	return dummy.Next
}
