

## 使用言語
- Go

## STEP1
実装後の後付けになりますが、コンピューターにお願いしたいことを整理しました。
- 空の箱を作る
- 紐でつながっている品物をばらして、箱の中に入れる
- 箱の中にある品物で、重複している品物は一個にする
- 紐でつなぎなおす

「紐でつなぎなおす」部分の実装がパッと思いつきませんでしたが、  
過去に二分木の実装で再帰関数を使用しているケースをみかけたことがあったのを思い出したため、  
過去の例を見ながら取り入れました。

```Go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "slices"

func deleteDuplicates(head *ListNode) *ListNode {
	var newList *ListNode
	node := head
	sortedList := make([]int, 0)

	for node != nil {
		sortedList = append(sortedList, node.Val)
		node = node.Next
	}

	noDuplicates := slices.Compact(sortedList)

	for _, v := range noDuplicates {
		newList = insert(newList, v)
	}

	return newList
}

func insert(l *ListNode, v int) *ListNode {
	if l == nil {
		return &ListNode{Val: v}
	}

	l.Next = insert(l.Next, v)

	return l
}
```

個人的には、「リストの順序は保証されている」という制約が、この関数と強い依存関係あることに引っかかりを覚えまして、この依存関係を分かるようにしないとバグの温床になりそうだなと思いました。  
公式パッケージ「slices」の関数「Compact」も、順番が保証されないと重複を削除しないので、`head`の順番が保証されないのなら動かないコードになってしまうと思いました。  
順序が保証されていないのであれば、「slices.Compact」の前に「slices.Sort」で箱の中身をソートかけると思います。

また、再帰関数は読み解くのに認知負荷が高いと思いました。（自分が例を見て理解するのに時間がかかった為）  


## STEP2
他解答を見て、下記がシンプルな方法だと感じました。  
自分なりの改良として、下記を行いました。
- `head`はソート済みであることを指す「sorted」に変更
- `else`の代わりに`continue`を使用

```Go
func deleteDuplicates(head *ListNode) *ListNode {
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
```

順序が絶対保証されているのであれば、この方法がシンプルで覚えやすいと思いました。  
「sorted.Next」がポインタであることが抜けてましたので、そこに気が付けば、headを返す理由の理解が進みました。


## STEP3
10分以内に3回何も見ないでコードを書けました。
```Go
func deleteDuplicates(head *ListNode) *ListNode {
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
```


## STEP4
テストを作成するのに時間がかかりました。  
特に、ランダムに数を生成しその値をソートしてListNodeの形にするところが、
難航しました。



- 時間計算量：O(n)
- 空間計算量：O(1)

**Constraints**
> - The number of nodes in the list is in the range [0, 300].
> - 100 <= Node.val <= 100
> - The list is guaranteed to be sorted in ascending order.



```
Benchmark

goos: windows
goarch: amd64
pkg: bench
cpu: Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz
BenchmarkDeleteDuplicates-8   	 5955951	       181.4 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	bench	1.633s
```

- BenchmarkDeleteDuplicates-8  
  ベンチマーク名 - ベンチマークのGOMAXPROCSの値


