package discovery

import (
	"container/heap"

	"github.com/livepeer/go-livepeer/net"
)

// A PriorityQueue implements heap.Interface and holds Items.
type priorityQueue []*suspension

// A suspension is the item we manage in the priority queue.
type suspension struct {
	orch    *net.OrchestratorInfo
	penalty int64
}

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].penalty > pq[j].penalty
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*suspension)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func newPriorityQueue() *priorityQueue {
	pq := &priorityQueue{}
	heap.Init(pq)
	return pq
}
