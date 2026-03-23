

## 使用言語
- Go

## STEP1
配列を別に作って、それを操作して重複が消えたものを改めてListNodeとして返却する方法を考えましたが、
まとまりがつかず5分が立ちましたので、新井氏の解説動画を見ました。

```Go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
  sentinel := &ListNode{Next: head}
  preNode := sentinel
  node := head

  for node != nil {
    if node != nil && node.Val == node.Next.Val {
        for node != nil && node.Val == node.Next.Val {
          node = node.Next
        }
      preNode.Next = node.Next
    } else {
      preNode = preNode.Next
    }
    node = node.Next
  }

  return sentinel.Next
}
```

Listの最初にダミーのnodeを置く発想はありませんでした。  
これを番兵（sentinel）と呼ばれることを初めて知りました。  
番兵はheadより一歩遅れて進むことが分かると、理解が進みました。

## STEP2
Listが[1, 1]のテストケースは落ちてしまったため、条件の見直しを行いました。  
また、ネストは浅い方が見やすいと感じたため、重複を見つける部分を移動しました。

```Go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
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
```


## STEP3
STEP2の記述と変わりませんが、3回とも10分以内でエラーなく通りました。  
よくある味見をしすぎて味が分からなくなるではありませんが、一度納得してしまうと視野が狭くなりそうな気がしましたので、積極的に他のコードを見る必要があると感じました。

```Go
func deleteDuplicates(head *ListNode) *ListNode {
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
```

## STEP4

- 時間計算量：O(n)
- 空間計算量：O(1)

**Constraints**
> The number of nodes in the list is in the range [0, 300].  
> -100 <= Node.val <= 100  
> The list is guaranteed to be sorted in ascending order.


```
Benchmark

goos: windows
goarch: amd64
pkg: bench/82-remove-duplicates-from-sorted-list-II
cpu: Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz
BenchmarkDeleteDuplicates-8   	 9424736	       119.8 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	bench/82-remove-duplicates-from-sorted-list-II	1.650s
```




