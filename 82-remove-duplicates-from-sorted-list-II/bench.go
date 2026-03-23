package bench

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	sentinel := &ListNode{Next: head}
	preNode := sentinel
	node := head

	for node != nil && node.Next != nil {
		if node.Val != node.Next.Val {
			preNode = preNode.Next
			node = node.Next
			continue
		}
		for node != nil && node.Next != nil && node.Val == node.Next.Val {
			node = node.Next
		}
		preNode.Next = node.Next
		node = node.Next
	}
	return sentinel.Next
}

func deleteDupulicates(head *ListNode) *ListNode {
	sentinel := &ListNode{Next: head}
	preNode := sentinel
	node := head

	for node != nil && node.Next != nil {
		if node.Val != node.Next.Val {
			preNode = preNode.Next
			node = node.Next
			continue
		}
		for node != nil && node.Next != nil && node.Val == node.Next.Val {
			node = node.Next
		}
		preNode.Next = node.Next
		node = node.Next
	}
	return sentinel.Next
}
