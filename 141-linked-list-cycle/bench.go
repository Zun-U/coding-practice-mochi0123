package bench

type ListNode struct {
	Val  int
	Next *ListNode
}

// フロイドの循環検出法
func HasCyclePointer(head *ListNode) bool {

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {

		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}

// map(set)による検出法
func HasCycleMap(head *ListNode) bool {

	nodes := make(map[*ListNode]bool)

	for head != nil {

		// カンマokイディオムを使用する
		// 「ok」以外の変数名も可能だが、Goのイディオムに従う
		_, ok := nodes[head]
		if ok {
			return true
		}

		nodes[head] = true
		head = head.Next
	}

	return false
}
