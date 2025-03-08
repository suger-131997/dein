package priorityqueue

import (
	"container/heap"
	"github.com/suger-131997/dein/internal/provider"
)

type PriorityQueue []*provider.Provider

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Out().Less(pq[j].Out())
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*provider.Provider)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func NewPriorityQueue() *PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return &pq
}

func Push(pq *PriorityQueue, item *provider.Provider) {
	heap.Push(pq, item)
}

func Pop(pq *PriorityQueue) *provider.Provider {
	item := heap.Pop(pq).(*provider.Provider)
	return item
}
