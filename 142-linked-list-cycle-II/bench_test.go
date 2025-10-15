package bench

import (
	"math/rand"
	"testing"
)

var blackhole *ListNode

const (
	MAXVALUE     = 10000
	MINVALUE     = -10000
	HEADPOSITION = 0
)

func makeCycleList(values []int, pos int) *ListNode {
	size := len(values)
	if size == 0 || pos < -1 {
		return nil
	}

	nodes := make([]*ListNode, size)

	for i, v := range values {
		nodes[i] = &ListNode{Val: v}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	if pos >= 0 && pos < size {
		nodes[size-1].Next = nodes[pos]
	}

	return nodes[HEADPOSITION]
}

func makeList(size, pos int) *ListNode {
	values := make([]int, size)

	for i := 0; i < size; i++ {
		values[i] = rand.Intn(MAXVALUE - MINVALUE + 1) + MINVALUE
	}

	return makeCycleList(values, pos)
}

func BenchmarkDetectCycleMap(b *testing.B) {

	nodes := makeList(10000, 2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := DetectCycleMap(nodes)
		blackhole = result
	}
}
