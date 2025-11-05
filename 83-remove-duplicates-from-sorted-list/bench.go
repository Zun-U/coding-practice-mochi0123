package bench

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	sorted := head

	for sorted != nil && sorted.Next != nil {
		if sorted.Val == sorted.Next.Val {
			sorted.Next = sorted.Next.Next
			continue
		}
		sorted = sorted.Next
	}

	return head
}
