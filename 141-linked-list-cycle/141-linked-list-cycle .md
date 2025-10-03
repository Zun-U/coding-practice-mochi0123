

## 使用言語
- Go

今後業務で使用する予定がある為、勉強も兼ねてGoを選択しました。

## STEP1
英文をそのまま読んでみましたが、問題文の意味が分かりませんでした。
日本語に機械翻訳しましたが、さほど理解が進まなかったので新井浩平氏の解説動画を視聴しました。
解き方は理解できましたが、自分はこの解法は所見では思いつかないと思います。
（フロイドの循環検出法というのを後で知りました）
また、`pos`が何を意味するのか、最初は全然わかりませんでしたが、
`pos`は`head`の終端にある要素の`Next`の位置の説明（ただの情報）であることは、あとから理解できました。

一回目は解説の通りに記述し、すらすらと記述できるまで解きました。

```Go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	fast := head
	slow := head

	for fast != nil {
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
```

何も見ないで記述できるようになることで、解法を理解したかどうか、セルフチェックが行えたと思います。


## STEP2
LeetCodeの他の解答を見つつ、変更した点は下記の通りです。
- 条件を一つにまとめる
- ガード節の削除
  - 理由としてはfor文の条件と同様なことを行っている
  - ループに条件を付けたい気持ちがある

```Go
func hasCycle(head *ListNode) bool {

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
```

別の解法だとmap(set)を使っているパターンを見かけました。

```Go
func hasCycle(head *ListNode) bool {

	nodes := make(map[*ListNode]bool, 0)

	for head != nil && head.Next != nil {

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

// 時間計算量：O(n)
// 空間計算量：O(n)
```


こちらは思いつきたかったな、と思いました。
ヘンゼルとグレーテルのようにパンくずを落とすような解法（道に迷わないように目印をする）は個人的に身近に感じますが、
一度フロイドの循環検出法を知ると、そちらの解法がすんなり入ってくるのは不思議でした。




## STEP3
どちらの解法も、10分以内に3回何も見ないでコードを書けました。  
自然に覚えられたと思います。


- フロイドの循環検出法
```go
func hasCycle(head *ListNode) bool {

	fast := head
	slow := head

	if head != nil {

		fast := fast.Next.Next
		slow := slow.Next

		if fast == slow {
			return ture
		}
	}

	return false
}
```

- map(set)を使用した解法
```go
func hasCycle(head *ListNode) bool {

	nodes := make(map[*ListNode]bool)

	if head != nil {

		_, ok := nodes[head]
		if ok {
			return true
		}

		nodes[head] = true
		head := head.Next
	}

	return false
}
```


## STEP4
オリジナルの手順として、ベンチマークの計測を行いました。
時間計算量、空間計算量の算出方法も分からなかったので、計算方法を勉強するのとあわせて実際にいくら時間がかかって、どれくらいメモリを使用するのか知りたかったのがモチベーションになります。

LeetCodeに`Constraints`の項目があったので、その項目に合わせて、node数1万のテストケースを作成します。

> - The number of the nodes in the list is in the range `[0, 10^4]`.
> - `-10^5 <= Node.val <= 10^5`
> - `pos` is `-1` or a **valid index** in the linked-list.





### フロイドの循環検出法
- 時間計算量：O(n)
- 空間計算量：O(1)
```
goos: windows
goarch: amd64
pkg: bench
cpu: Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz
BenchmarkHasCycle-8        50950             24817 ns/op               6 B/op          0 allocs/op
PASS
ok      bench   1.872s
```


### map(set)
- 時間計算量：O(n)
- 空間計算量：O(n)
```
goos: windows
goarch: amd64
pkg: bench
cpu: Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz
BenchmarkHasCycle-8         1382            846686 ns/op          591722 B/op         86 allocs/op
PASS
ok      bench   1.646s
```

フロイドの循環検出法が実行時間、メモリ共にパフォーマンスが良い結果となりました。
時間計算量はどちらもO(n)でしたが、差が出る結果となりました。
もっと大きいnのケースであれば、大体同じくらいの時間になるのでしょうか。

組み込みのテスト関数の内部は調べてみたいと思いました。


## STEP5
レビューをもとに修正してみました。


```go
func hasCycle(head *ListNode) bool {

	// あらかじめ要素の個数が分かっている場合は、サイズを指定する
	// 今回は要素の個数が分からないため、サイズを指定しない（サイズは0になる）
	nodes := make(map[*ListNode]bool)

	// headはあくまでnodeの先頭を意味するため、別に変数を作りそちらを動かす
	visited := head

	// ループの条件として、head.Next != nil は必須ではないため、省略
	for head != nil {

		_, ok := nodes[visited]
		if ok {
			return true
		}

		nodes[visited] = true
		visited = head.Next
	}

	return false
}
```