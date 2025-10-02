package bench

import (
	"testing"
	"math/rand"
)

// 実行結果をエスケープするための変数
// コンパイラが気を利かせてデットコードを削除（最適化）するのを防ぐ
var blackhole bool

const (
	MINVALUE int = -10000
	MAXVALUE int = 10000
)

func makeCycleList(values []int, pos int) *ListNode{

	if len(values) == 0 || -1 > pos {
		return nil
	}

	nodes := make([]*ListNode, len(values))

	for i, v := range values {
		nodes[i] = &ListNode{
			Val: v,
		}
		if i > 0 {
			// 一個前のNodeのNextに、現在のNodeを紐づける
			nodes[i-1].Next = nodes[i]
		}
	}

	if pos >= 0 && pos < len(values) {
		// 最後のNodeのNextに、posのNodeを紐づける
		nodes[len(values)-1].Next = nodes[pos]
	}

	return  nodes[0]
}

func makeList(size, pos int) *ListNode {

	values := make([]int, size)

	for i := 0; i < size; i++ {
		// -10^5 <= value <= 10^5 の範囲でランダムな値のsliceを作成
		values[i] = rand.Intn(MAXVALUE - MINVALUE + 1) + MINVALUE
	}

	return makeCycleList(values, pos)
}

func BenchmarkHasCycle(b *testing.B) {

	// 10^4のListでテスト
	nodes := makeList(10000, 1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := HasCyclePointer(nodes)
		blackhole = result
	}

}
