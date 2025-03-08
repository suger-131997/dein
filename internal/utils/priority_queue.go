package utils

import (
	"container/heap"
)

type PriorityQueue[T any] struct {
	queue []T
	less  func(i, j T) bool
}

func (pq *PriorityQueue[T]) Len() int { return len(pq.queue) }

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return pq.less(pq.queue[i], pq.queue[j])
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(T)
	pq.queue = append(pq.queue, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := pq.queue
	n := len(old)
	item := old[n-1]
	pq.queue = old[0 : n-1]
	return item
}

func NewPriorityQueue[T any](less func(i, j T) bool) *PriorityQueue[T] {
	queue := make([]T, 0)
	pq := &PriorityQueue[T]{
		queue: queue,
		less:  less,
	}
	heap.Init(pq)
	return pq
}

func Push[T any](pq *PriorityQueue[T], item T) {
	heap.Push(pq, item)
}

func Pop[T any](pq *PriorityQueue[T]) T {
	item := heap.Pop(pq).(T)
	return item
}
