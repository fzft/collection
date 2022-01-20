package collection

import "fmt"

const (
	DEFAULT_CAPACITY = 10
)

// ArrayList resizeable-array implementation of the List interface
type ArrayList struct {
	elementData []interface{}
	size        int
}

// NewArrayList constructs an empty list with the specified initial capacity
func NewArrayList(initialCapacity int) *ArrayList {
	var initCap int
	if initialCapacity < 0 {
		panic(fmt.Sprintf("Illegal Capacity: %d", initialCapacity))
	} else if initialCapacity == 0 {
		initCap = DEFAULT_CAPACITY
	} else {
		initCap = initialCapacity
	}

	return &ArrayList{elementData: make([]interface{}, initCap)}
}

func NewArrayListWithSlice(s []interface{}) *ArrayList {
	l := new(ArrayList)
	l.size = len(s)
	l.elementData = make([]interface{}, l.size)
	copy(l.elementData, s)
	return l
}

// Get return the element at the specified position in this list
func (l *ArrayList) Get(index int) interface{} {
	return l.elementData[index]
}

func (l *ArrayList) GetSize() int {
	return l.size
}

func (l *ArrayList) IsEmpty() bool {
	return l.size == 0
}

func (l *ArrayList) Contains(o interface{}) bool {
	return l.indexOf(o) >= 0
}

func (l *ArrayList) Remove(o interface{}) bool {
	for i, v := range l.elementData {
		if v == o {
			l.fastRemove(i)
			return true
		}
	}
	return false
}

func (l *ArrayList) RemoveWithIndex(index int) interface{} {
	l.rangeCheckForAdd(index)
	oldValue := l.elementData[index]
	l.fastRemove(index)
	return oldValue
}

// AddAll the implementation use two allocate
func (l *ArrayList) AddAll(index int, c Collection) bool {
	a := c.Iterator()
	numNew := c.GetSize()

	tmpData := make([]interface{}, l.size)
	copy(tmpData, l.elementData)

	l.ensureCapacityInternal(l.size + numNew)
	copy(l.elementData[index:], a[:])
	copy(l.elementData[index+numNew:], tmpData[index:])

	l.size += numNew
	return numNew != 0
}

func (l *ArrayList) RemoveAll(c Collection) bool {
	return l.batchRemove(c, false)
}

func (l *ArrayList) Iterator() []interface{} {
	return l.elementData
}

func (l *ArrayList) Clear() {
	l.elementData = make([]interface{}, cap(l.elementData))
	l.size = 0
}

// Add appends the specified element to the end of this list
func (l *ArrayList) Add(e interface{}) bool {
	l.ensureCapacityInternal(l.size + 1)
	l.elementData[l.size] = e
	l.size++
	return true
}

// RetainAll retains only the elements in this list that are contained in the specified collection
func (l *ArrayList) RetainAll(c Collection) bool {
	return l.batchRemove(c, true)
}

// Insert the specified element at the specified position in this list
func (l *ArrayList) Insert(index int, e interface{}) {
	l.rangeCheckForAdd(index)
	l.ensureCapacityInternal(l.size + 1)
	copy(l.elementData[index+1:], l.elementData[index:])
	l.elementData[index] = e
	l.size++
}

// fastRemove ... skip bound checking
func (l *ArrayList) fastRemove(index int) {
	copy(l.elementData[index:], l.elementData[index+1:])
	l.elementData[l.size-1] = nil
	l.elementData = l.elementData[:l.size-1]
	l.size--
}

// trimToSize trim the capacity of this ArrayList instance to be the list's current size.
// this operation to minimize the storage of an ArrayList instance
func (l *ArrayList) trimToSize() {
	if l.size < cap(l.elementData) {
		if l.size == 0 {
			l.elementData = make([]interface{}, 0)
		} else {
			newElementData := make([]interface{}, l.size)
			copy(newElementData, l.elementData)
			l.elementData = newElementData
		}
	}
}

// indexOf returns the index of the first occurrence of the specified element in this list, or -1 if this list does not
// contain the element
func (l *ArrayList) indexOf(o interface{}) int {
	for i, v := range l.elementData {
		if v == o {
			return i
		}
	}
	return -1
}

// lastIndexOf returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element
func (l *ArrayList) lastIndexOf(o interface{}) int {
	for i := l.size - 1; i >= 0; i-- {
		if l.elementData[i] == o {
			return i
		}
	}
	return -1
}

func (l *ArrayList) ensureCapacityInternal(minCapacity int) {
	if cap(l.elementData) == 0 {
		minCapacity = Max(DEFAULT_CAPACITY, minCapacity)
	}
	l.ensureExplicitCapacity(minCapacity)
}

// grow increase the capacity of the array
func (l *ArrayList) grow(minCap int) {
	oldCap := cap(l.elementData)
	newCap := oldCap + oldCap>>1
	if newCap-minCap < 0 {
		newCap = minCap
	}
	newElementData := make([]interface{}, newCap)
	copy(newElementData, l.elementData)
	l.elementData = newElementData
}

func (l *ArrayList) rangeCheckForAdd(index int) {
	if index > l.size || index < 0 {
		panic(fmt.Sprintf("%d index out of Bound", index))
	}
}

func (l *ArrayList) ensureExplicitCapacity(minCapacity int) {
	if minCapacity-cap(l.elementData) > 0 {
		l.grow(minCapacity)
	}
}

func (l *ArrayList) batchRemove(c Collection, complement bool) bool {
	var (
		r        int
		w        int
		modified bool
	)
	for ; r < l.size; r++ {
		if c.Contains(l.elementData[r]) == complement {
			l.elementData[w] = l.elementData[r]
			w++
		}
	}
	if w != l.size {
		for i := w; i < l.size; i++ {
			l.elementData[i] = nil
		}
		l.size = w
		modified = true
	}
	return modified
}
