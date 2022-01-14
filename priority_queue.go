package collection

type PriorityQueue struct {
	*AbstractQueue
	queue      []interface{}
	size       int
	comparator Comparator
	// modCount the number of times this priority queue has been structurally modified
	modCount int
}

func NewPriorityQueue(initCap int, comparator Comparator) *PriorityQueue {
	pq := new(PriorityQueue)
	pq.AbstractQueue = &AbstractQueue{}
	pq.comparator = comparator
	pq.queue = make([]interface{}, initCap)
	return pq
}

func (p *PriorityQueue) GetSize() int {
	return p.size
}

func (p *PriorityQueue) Contains(o interface{}) bool {
	panic("implement me")
}

func (p *PriorityQueue) Add(e interface{}) bool {
	if e == nil {
		panic("null pointer exception")
	}
	p.modCount++
	i := p.size
	if i >= cap(p.queue) {
		p.grow(i + 1)
	}
	p.size = i + 1
	if i == 0 {
		p.queue[0] = e
	} else {
		p.siftUp(i, e)
	}
	return true

}

func (p *PriorityQueue) Remove(o interface{}) bool {
	for i, e := range p.Iterator() {
		if o == e {
			p.removeAt(i)
			return true
		}
	}
	return false
}

func (p *PriorityQueue) Iterator() []interface{} {
	return p.queue
}

func (p *PriorityQueue) Clear() {
	p.modCount++
	p.queue = make([]interface{}, cap(p.queue))
	p.size = 0
}

func (p *PriorityQueue) Poll() interface{} {
	if p.size == 0 {
		return nil
	}
	p.size--
	p.modCount++
	result := p.queue[0]
	x := p.queue[p.size]
	p.queue[p.size] = nil
	if p.size != 0 {
		p.siftDown(0, x)
	}
	return result
}

func (p *PriorityQueue) Peek() interface{} {
	if p.size == 0 {
		return nil
	}
	return p.queue[0]
}

// siftDown inserts item x at position k, maintaining heap invariant by demoting x down the tree repeatedly until
// it is less than or equal to its children or is a leaf
func (p *PriorityQueue) siftDown(k int, x interface{}) {
	if p.comparator != nil {
		p.siftDownUsingComparator(k, x)
	} else {
		panic("comparator is nil")
	}
}

func (p *PriorityQueue) siftDownUsingComparator(k int, x interface{}) {
	half := p.size >> 1
	for k < half {
		child := (k << 1) + 1
		c := p.queue[child]
		right := child + 1
		if right < p.size && p.comparator(c, p.queue[right]) > 0 {
			child = right
			c = p.queue[child]
		}
		if p.comparator(x, c) <= 0 {
			break
		}
		p.queue[k] = c
		k = child
	}
	p.queue[k] = x
}

// grow increase the capacity of the array
func (p *PriorityQueue) grow(minCap int) {
	var newCap int
	oldCap := cap(p.queue)
	// double size if small ; else grow by 50%
	if oldCap < 64 {
		newCap = oldCap << 1 + 2
	} else {
		newCap = oldCap + oldCap>>1
	}
	newQueue := make([]interface{}, newCap)
	copy(newQueue, p.queue)
	p.queue = newQueue
}

// siftUp insert item x at position k, maintaining heap invariant by promoting x up the tree until
// it is greater than or equal to its parent, or is the root.
func (p *PriorityQueue) siftUp(k int, e interface{}) {
	if p.comparator != nil {
		p.siftUpUsingComparator(k, e)
	} else {
		panic("comparator is nil")
	}
}

func (p *PriorityQueue) siftUpUsingComparator(k int, x interface{}) {
	for k > 0 {
		parent := (k - 1) >> 1
		e := p.queue[parent]
		if p.comparator(x, e) >= 0 {
			break
		}
		p.queue[k] = e
		k = parent
	}
	p.queue[k] = x
}

// removeAt removes the ith element from the queue. in order to maintain the heap invariant, it must swap a later element of the list
// with one earlier than i
func (p *PriorityQueue) removeAt(i int) interface{} {
	p.modCount++
	p.size--
	s := p.size

	if s == i { // removed last element
		p.queue[i] = nil
	} else {
		moved := p.queue[s]
		p.queue[s] = nil
		p.siftDown(i, moved)
		if p.queue[i] == moved {
			p.siftUp(i, moved)
			if p.queue[i] != moved {
				return moved
			}
		}
	}
	return nil
}
