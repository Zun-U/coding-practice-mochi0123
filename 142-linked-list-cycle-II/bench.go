package bench

type ListNode struct {
	Val  int
	Next *ListNode
}

func DetectCycleMap(head *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)
	current := head

	for current != nil {
		_, ok := visited[current]
		if ok {
			return current
		}

		visited[current] = true
		current = current.Next
	}

	return nil
}

func DetectCycleFloyd(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			slow = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}

	return nil
}
