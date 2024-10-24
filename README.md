# heap

heap是对标准库`container/heap`的简单封装。使堆结构更简单易用。



## 示例

普通小顶堆：

```go
package main

import (
	"fmt"

	"github.com/eachain/heaps"
)

func main() {
	h := heaps.NewOrderedHeap[int]()
	h.Push(2)
	h.Push(1)
	h.Push(5)
	h.Push(3)
	fmt.Printf("minimum: %d\n", h.Top())
	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}
```

优先级队列：

```go
package main

import (
	"fmt"

	"github.com/eachain/heaps"
)

func main() {
	pq := heaps.NewPriorityQueue[string](func(a, b int) bool {
		return a > b
	})
	pq.Push("banana", 3)
	pq.Push("apple", 2)
	pq.Push("pear", 4)
	pq.Push("orange", 1)
	pq.Update("orange", 5)

	for pq.Len() > 0 {
		item, priority := pq.Pop()
		fmt.Printf("%.2d:%s ", priority, item)
	}
}
```




