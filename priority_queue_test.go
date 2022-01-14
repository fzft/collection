package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type PriorityQueueTest struct {
	pq *PriorityQueue
}

func PriorityQueueTestSetup(tb testing.TB) (func(tb testing.TB), PriorityQueueTest) {
	pt := PriorityQueueTest{}
	pt.pq = NewPriorityQueue(5, func(t, u interface{}) int {
		val1 := t.(int)
		val2 := u.(int)
		if val1 > val2 {
			return 1
		} else if val1 == val2 {
			return 0
		} else {
			return -1
		}
	})

	return func(tb testing.TB) {
		tb.Log("PriorityQueueTestSetup teardown")
	}, pt
}

func TestPriorityQueue_Add(t *testing.T) {
	teardownTest, pt := PriorityQueueTestSetup(t)
	defer teardownTest(t)

	// when
	pt.pq.Add(3)
	pt.pq.Add(2)
	pt.pq.Add(5)
	pt.pq.Add(4)

	// then
	assert.Equal(t, 4, pt.pq.GetSize())
}

func TestPriorityQueue_Peek(t *testing.T) {
	teardownTest, pt := PriorityQueueTestSetup(t)
	defer teardownTest(t)

	// when
	pt.pq.Add(3)
	pt.pq.Add(2)
	pt.pq.Add(5)
	pt.pq.Add(4)
	pt.pq.Add(1)

	// then
	assert.Equal(t, 1, pt.pq.Peek())
}

func TestPriorityQueue_Poll(t *testing.T) {
	teardownTest, pt := PriorityQueueTestSetup(t)
	defer teardownTest(t)

	// when
	pt.pq.Add(3)
	pt.pq.Add(2)
	pt.pq.Add(5)
	pt.pq.Add(4)
	pt.pq.Add(1)

	// then
	expectedSeq := []int{1, 2, 3, 4, 5}

	for i := 0; i < 5; i++ {
		assert.Equal(t, expectedSeq[i], pt.pq.Poll())
	}

}

func TestPriorityQueue_grow(t *testing.T) {
	teardownTest, pt := PriorityQueueTestSetup(t)
	defer teardownTest(t)

	// when
	pt.pq.Add(3)
	pt.pq.Add(2)
	pt.pq.Add(5)
	pt.pq.Add(4)
	pt.pq.Add(1)

	assert.Equal(t, 5, cap(pt.pq.queue))

	//
	pt.pq.Add(3)

	// then
	assert.Equal(t, 12, cap(pt.pq.queue))
}

func TestPriorityQueue_Remove(t *testing.T) {
	teardownTest, pt := PriorityQueueTestSetup(t)
	defer teardownTest(t)

	// when
	pt.pq.Add(3)
	pt.pq.Add(2)
	pt.pq.Add(5)
	pt.pq.Add(4)
	pt.pq.Add(1)

	pt.pq.Remove(3)

	// then
	expectedSeq := []int{1, 2, 4, 5}

	for i := 0; i < 4; i++ {
		assert.Equal(t, expectedSeq[i], pt.pq.Poll())
	}
}
