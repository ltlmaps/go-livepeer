package discovery

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	assert := assert.New(t)
	pq := newPriorityQueue()
	susp0 := &suspension{penalty: 4}
	heap.Push(pq, susp0)
	assert.Equal(pq.Len(), 1)
	assert.Equal(pq.Pop().(*suspension), susp0)
	assert.Equal(pq.Len(), 0)

	susp1 := &suspension{penalty: 6}
	heap.Push(pq, susp0)
	assert.Equal(pq.Len(), 1)
	heap.Push(pq, susp1)
	assert.Equal(pq.Len(), 2)
	assert.Equal(pq.Pop().(*suspension), susp0)
	assert.Equal(pq.Len(), 1)
	assert.Equal(pq.Pop().(*suspension), susp1)
	assert.Equal(pq.Len(), 0)

	susp2 := &suspension{penalty: 2}
	heap.Push(pq, susp0)
	assert.Equal(pq.Len(), 1)
	heap.Push(pq, susp1)
	assert.Equal(pq.Len(), 2)
	heap.Push(pq, susp2)
	assert.Equal(pq.Len(), 3)
	assert.Equal(pq.Pop().(*suspension), susp2)
	assert.Equal(pq.Len(), 2)
	assert.Equal(pq.Pop().(*suspension), susp0)
	assert.Equal(pq.Len(), 1)
	assert.Equal(pq.Pop().(*suspension), susp1)
	assert.Equal(pq.Len(), 0)
}
