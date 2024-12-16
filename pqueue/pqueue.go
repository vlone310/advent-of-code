package pqueue

import "container/heap"

// PriorityQueue is a generic priority queue that works with any type T
// that has a priority (defined by a lessFunc comparator).
type PriorityQueue[T any] struct {
	items    []T
	lessFunc func(a, b T) bool
}

// New creates a new priority queue with a given comparator function.
func New[T any](lessFunc func(a, b T) bool) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{lessFunc: lessFunc}
	heap.Init(pq)
	return pq
}

func (pq PriorityQueue[T]) Len() int { return len(pq.items) }
func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq.lessFunc(pq.items[i], pq.items[j])
}
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	item := x.(T)
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue[T]) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[0 : n-1]
	return item
}

// PushItem adds an item to the priority queue.
func (pq *PriorityQueue[T]) PushItem(item T) {
	heap.Push(pq, item)
}

// PopItem removes and returns the item with the highest priority.
func (pq *PriorityQueue[T]) PopItem() T {
	return heap.Pop(pq).(T)
}
