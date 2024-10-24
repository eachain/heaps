package heaps

import (
	"math/rand"
	"testing"
)

func TestOrder(t *testing.T) {
	const N = 1000
	values := make([]int, N)
	for i := range values {
		values[i] = rand.Int()
	}

	h := NewOrderedHeap[int]()
	for _, v := range values {
		h.Push(v)
	}

	prev := h.Pop()
	for h.Len() > 0 {
		val := h.Pop()
		if prev > val {
			t.Fatalf("heap pop prev %v > next %v", prev, val)
		}
		prev = val
	}
}

func TestTop(t *testing.T) {
	h := NewOrderedHeap[int]()
	for i := 1000; i > 0; i-- {
		h.Push(i)
	}
	top := h.Top()
	if top != 1 {
		t.Fatalf("heap top: %v", top)
	}
}

func TestCount(t *testing.T) {
	const N = 1000
	values := make([]int, N)
	count := make(map[int]int)
	for i := range values {
		values[i] = rand.Int()
		count[values[i]]++
	}

	h := NewOrderedHeap[int]()
	for _, v := range values {
		h.Push(v)
	}

	for v, n := range count {
		m := h.Count(v)
		if m != n {
			t.Fatalf("heap count %v: %v, actually %v", v, m, n)
		}
	}
}

func TestRemove(t *testing.T) {
	const N = 1000
	values := make([]int, N)
	count := make(map[int]int)
	for i := range values {
		values[i] = rand.Int()
		count[values[i]]++
	}

	h := NewOrderedHeap[int]()
	for _, v := range values {
		h.Push(v)
	}

	for v, n := range count {
		m := h.Remove(v, n)
		if m != n {
			t.Fatalf("heap remove %v: %v, actually %v", v, m, n)
		}
	}

	if n := h.Len(); n != 0 {
		t.Fatalf("heap remove result len: %v", n)
	}
}
