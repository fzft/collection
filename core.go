package collection

// Collection the root interface in the collection hierarchy
type Collection interface {
	// GetSize returns the number of elements in this collection
	GetSize() int

	// IsEmpty return true if this collection contains no elements
	IsEmpty() bool

	// Contains true if this collection contains the specified element
	Contains(o interface{}) bool

	// Add ...
	Add(e interface{}) bool

	// Remove a single instance of the specified element from this collection, if it is present.
	Remove(o interface{}) bool

	// Clear removes all the elements from this collection
	Clear()

	// Iterator return the elements contained in this collection
	Iterator() []interface{}
}

// AbstractCollection ...
type AbstractCollection struct {

}

func (a *AbstractCollection) Iterator() []interface{} {
	panic("implement me")
}

func (a *AbstractCollection) GetSize() int {
	panic("implement me")
}

func (a *AbstractCollection) IsEmpty() bool {
	return a.GetSize() == 0
}

func (a *AbstractCollection) Contains(o interface{}) bool {
	panic("implement me")
}

func (a *AbstractCollection) Add(e interface{}) bool {
	panic("implement me")
}

func (a *AbstractCollection) Remove(o interface{}) bool {
	panic("implement me")
}

func (a *AbstractCollection) Clear() {
	panic("implement me")
}

func NewAbstractCollection() *AbstractCollection {
	return &AbstractCollection{}
}



// Queue a collection designed for holding elements prior to processing
type Queue interface {
	Collection

	// Poll retrieves and removes the head of this queue, or return nil if this queue is empty
	Poll() interface{}

	// Peek retrieves but not remove, the head of this queue, or return nil if this queue is empty
	Peek() interface{}
}

type AbstractQueue struct {
	*AbstractCollection
}


