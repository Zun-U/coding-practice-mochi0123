package bench

import (
	"math/rand"
	"slices"
	"testing"
	"time"
)

var blackhole *ListNode

const (
	MAXVALUE     = 100
	MINVALUE     = -100
	MAXRANGE     = 300
	HEADPOSITION = 0
)

func makeList(size int) *ListNode {
	var nodeValue int
	nodes := make([]*ListNode, size)
	list  := make([]int, 0)

	// シード値の設定
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		// 半開区間[0,201)の乱数(0~200の間) - 100
		nodeValue = rand.Intn(MAXVALUE - MINVALUE + 1) + MINVALUE
		list      = append(list, nodeValue)
	}

	slices.Sort(list)

	for i, v := range list {
		nodes[i] = &ListNode{Val: v}
		if 0 < i {
			// 一個前のnodeのNextは、現在のnode
			nodes[i-1].Next = nodes[i]
		}
	}

	return nodes[HEADPOSITION]
}

func BenchmarkDeleteDuplicates(b *testing.B) {
	nodes := makeList(MAXRANGE)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := DeleteDuplicates(nodes)
		blackhole = result
	}
}