package heaps

import (
	"cmp"
	"container/heap"
	"unsafe"
)

// LessFunc compare a and b,
// it should returns true on a < b,
// and false on a >= b.
type LessFunc[T any] func(a, b T) bool

type Heap[T any] struct {
	addr   hAddr[T]
	values []T
	less   LessFunc[T]
}

func NewHeap[T any](less LessFunc[T]) *Heap[T] {
	return &Heap[T]{less: less}
}

func NewOrderedHeap[T cmp.Ordered]() *Heap[T] {
	return NewHeap(orderedLess[T])
}

func orderedLess[T cmp.Ordered](a, b T) bool {
	return a < b
}

func (h *Heap[T]) Push(v T) {
	heap.Push(&h.addr, v)
}

func (h *Heap[T]) Pop() T {
	return heap.Pop(&h.addr).(T)
}

func (h *Heap[T]) Len() int {
	return len(h.values)
}

func (h *Heap[T]) Top() T {
	return h.values[0]
}

func (h *Heap[T]) equal(a, b T) bool {
	return !h.less(a, b) && !h.less(b, a)
}

func (h *Heap[T]) scanUntil(i int, v T, f func(int)) {
	if i >= len(h.values) {
		return
	}
	if h.less(v, h.values[i]) {
		return
	}
	f(i)
	h.scanUntil(i*2+1, v, f) // left
	h.scanUntil(i*2+2, v, f) // right
}

// Count returns the count of elements equal v.
func (h *Heap[T]) Count(v T) int {
	n := 0
	h.scanUntil(0, v, func(i int) {
		if h.equal(h.values[i], v) {
			n++
		}
	})
	return n
}

// Count returns the count of elements equal v removed.
// The param n points how many elements to remove, -1 named all.
func (h *Heap[T]) Remove(v T, n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		n = len(h.values)
	}
	removed := 0
	h.scanUntil(0, v, func(i int) {
		if removed >= n {
			return
		}
		for i < len(h.values) && h.equal(h.values[i], v) {
			heap.Remove(&h.addr, i)
			removed++
		}
	})
	return removed
}

type hAddr[T any] struct{}

func (hAddr *hAddr[T]) heap() *Heap[T] {
	return (*Heap[T])(unsafe.Pointer(hAddr))
}

func (hAddr *hAddr[T]) Len() int {
	return len(hAddr.heap().values)
}

func (hAddr *hAddr[T]) Less(i, j int) bool {
	h := hAddr.heap()
	return h.less(h.values[i], h.values[j])
}

func (hAddr *hAddr[T]) Swap(i, j int) {
	h := hAddr.heap()
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (hAddr *hAddr[T]) Push(x any) {
	h := hAddr.heap()
	h.values = append(h.values, x.(T))
}

func (hAddr *hAddr[T]) Pop() any {
	h := hAddr.heap()
	n := len(h.values)
	x := h.values[n-1]
	h.values = h.values[0 : n-1]
	return x
}
