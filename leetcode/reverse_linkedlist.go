package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	previous, current := (*ListNode)(nil), head

	for current != nil {
		nxt := current.Next
		current.Next = previous
		previous = current
		current = nxt
	}

	return previous
}
