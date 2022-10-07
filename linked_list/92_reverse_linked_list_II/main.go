package _2_reverse_linked_list_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	var prev *ListNode
	for i := 0; i < left; i++ {
		prev = head
		head = head.Next
	}
	var mid *ListNode
	next := head
	for j := left; j <= right; j++ {
		if head == nil {
			break
		}
		temp := head.Next
		head.Next = mid
		mid = head
		head = temp
	}
	prev.Next = mid
	next.Next = head
	return dummy.Next
}
