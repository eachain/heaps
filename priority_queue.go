package heaps

import (
	"cmp"
	"container/heap"
	"unsafe"
)

type PriorityQueue[E comparable, P any] struct {
	addr     pqAddr[E, P]
	items    []E
	priority []P
	index    map[E]int
	less     LessFunc[P]
}

func NewPriorityQueue[E comparable, P any](less LessFunc[P]) *PriorityQueue[E, P] {
	return &PriorityQueue[E, P]{less: less}
}

func NewOrderedPriorityQueue[E comparable, P cmp.Ordered]() *PriorityQueue[E, P] {
	return &PriorityQueue[E, P]{less: orderedLess[P]}
}

func (pq *PriorityQueue[E, P]) Push(item E, priority P) {
	pq.priority = append(pq.priority, priority)
	heap.Push(&pq.addr, item)
}

func (pq *PriorityQueue[E, P]) Update(item E, priority P) bool {
	i, ok := pq.index[item]
	if !ok {
		return false
	}
	pq.priority[i] = priority
	heap.Fix(&pq.addr, i)
	return true
}

func (pq *PriorityQueue[E, P]) Pop() (E, P) {
	priority := pq.priority[0]
	return heap.Pop(&pq.addr).(E), priority
}

func (pq *PriorityQueue[E, P]) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue[E, P]) Top() (E, P) {
	return pq.items[0], pq.priority[0]
}

func (pq *PriorityQueue[E, P]) PriorityOf(item E) (priority P, ok bool) {
	i, ok := pq.index[item]
	if !ok {
		return
	}
	return pq.priority[i], true
}

func (pq *PriorityQueue[E, P]) Remove(item E) (priority P, ok bool) {
	i, ok := pq.index[item]
	if !ok {
		return
	}
	priority = pq.priority[i]
	ok = true
	heap.Remove(&pq.addr, i)
	return
}

type pqAddr[E comparable, P any] struct{}

func (pqAddr *pqAddr[E, P]) queue() *PriorityQueue[E, P] {
	return (*PriorityQueue[E, P])(unsafe.Pointer(pqAddr))
}

func (pqAddr *pqAddr[E, P]) Len() int {
	return len(pqAddr.queue().items)
}

func (pqAddr *pqAddr[E, P]) Less(i, j int) bool {
	pq := pqAddr.queue()
	return pq.less(pq.priority[i], pq.priority[j])
}

func (pqAddr *pqAddr[E, P]) Swap(i, j int) {
	pq := pqAddr.queue()
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.priority[i], pq.priority[j] = pq.priority[j], pq.priority[i]
	pq.index[pq.items[i]] = i
	pq.index[pq.items[j]] = j
}

func (pqAddr *pqAddr[E, P]) Push(x any) {
	pq := pqAddr.queue()
	item := x.(E)
	if pq.index == nil {
		pq.index = make(map[E]int)
	}
	pq.index[item] = len(pq.items)
	pq.items = append(pq.items, item)
}

func (pqAddr *pqAddr[E, P]) Pop() any {
	pq := pqAddr.queue()
	n := len(pq.items)
	item := pq.items[n-1]
	pq.items = pq.items[0 : n-1]
	pq.priority = pq.priority[0 : n-1]
	delete(pq.index, item)
	return item
}
