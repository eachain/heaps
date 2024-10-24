package heaps

import "testing"

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue[string](func(a, b int) bool {
		return a > b
	})
	pq.Push("banana", 3)
	pq.Push("apple", 2)
	pq.Push("pear", 4)
	pq.Push("orange", 1)
	pq.Update("orange", 5)

	fruits := []string{"orange", "pear", "banana", "apple"}
	pris := []int{5, 4, 3, 2}

	for i, f := range fruits {
		fruit, pri := pq.Pop()
		if fruit != f || pri != pris[i] {
			t.Fatalf("first: %v %v", fruit, pri)
		}
	}

}
