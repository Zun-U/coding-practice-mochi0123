

## 使用言語
- Go

## STEP1
英文をそのまま読んでみました。
前回の[141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)と、同じような文章だったので、何とか意味は理解できました。
map(set)を使用した方法では返り値を変えるだけで、一旦通りました。

```Go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
  nodes := make(map[*ListNode]bool)

  for head != nil {
    _, ok := nodes[head]
    if ok {
      return head
    }

    nodes[head] = true
    head = head.Next
  }

  return nil
}

```

新井浩平氏の解説動画を見ると、初見では理解できませんでした。  
`fast`が2倍の速度で移動するのでサイクルを必ず1周以上すること、`slow`は絶対に1周しないことが分かると、理解が進みました。  

```go
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	if fast != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

```

何も見ずに繰り返すと、条件文の指定ミスや、代入ルールミス等でやり直しが何回か発生しました。  
基礎的なところで、変な思考の癖があると思いましたので、思考とコードが一致するまで続けました。


## STEP2
下記変更を行いました。
- nodes　→　visitedに変更
- head　→　currentに代入

```Go
func detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)
	current := head

	for current != nil  {
		_, ok := visited[current]
		if ok {
			return current
		}

		visited[current] = true
		current = current.Next
	}

	return nil
}
```
関数の引数名を変えたいところですが、もしこの関数の影響範囲が大きければ...と考えると、気軽には変えられないと思いました。  
最初にこの関数を実装するのであれば、説明用の変数を作成せずに、直接引数名を変えてると思います。



フロイドの循環検出法は下記修正を行いました。
- 主役をfastで統一する
- 不要な値チェックを省略する

for文が入れ子になってしまいますが、処理の流れは自然でシンプルに見えました。

```Go
func detectCycle(head *ListNode) *ListNode {
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
```


## STEP3
どちらの解法も、10分以内に3回何も見ないでコードを書けました。  

- map(set)を使用した解法
```go
func detectCycle(head *ListNode) *ListNode {
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
```


- フロイドの循環検出法
```go
func detectCycle(head *ListNode) *ListNode {
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
```


## STEP4
ベンチマークの計測です。  
※条件は前回の[141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)と同様。

**Constraints**
> - The number of the nodes in the list is in the range `[0, 10^4]`.
> - `-10^5 <= Node.val <= 10^5`
> - `pos` is `-1` or a **valid index** in the linked-list.

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

### フロイドの循環検出法
- 時間計算量：O(n)
- 空間計算量：O(1)
```
goos: windows
goarch: amd64
pkg: bench
cpu: Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz
BenchmarkDetectCycleMap-8          49768             23135 ns/op               0 B/op          0 allocs/op
PASS
ok      bench   1.809s
```

「141. Linked List Cycle」と似たような結果になりました。  
同じO(1)でも、ハッシュ値の計算が時間がかかる理由を調べましたが、処理ステップ数がかかる為、定数倍の影響が大きいと理解いたしました。

- 参考資料  
[GopherCon 2016: Inside the Map Implementation - Keith Randall](https://www.youtube.com/watch?v=Tl7mi9QmLns&list=PLnPR191a_BqY1D5n4An5gfQf-LhHclKlS&index=75)


## STEP5
いただいたレビューをもとに修正しました。

- map(set)を使用した解法
  - `current`を`node`に変数名変更
```go
func detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)
	node := head

	for node != nil {
		_, ok := visited[node]
		if ok {
			return node
		}

		visited[node] = true
		node = node.Next
	}

	return nil
}
```


- フロイドの循環検出法
  - リストの循環検出部分を関数化する
```go

func hasCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return fast
		}
	}

	return nil
}

func detectCycle(head *ListNode) *ListNode {
	meetingNode := hasCycle(head)
	if meetingNode == nil {
		return nil
	}

	fromHead    := head
	fromMeeting := meetingNode
	for fromHead != fromMeeting {
		fromHead    = fromHead.Next
		fromMeeting = fromMeeting.Next
	}
	return fromHead
}
```