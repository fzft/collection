package collection

// List an ordered collection. The user of this interface has precise control over where in the list each element is inserted
type List interface {
	Collection

	RemoveWithIndex(index int) interface{}

	// RemoveAll removes from this list all of its elements that are contained in the specified collection
	RemoveAll(c Collection) bool

	// AddAll inserts all the elements in the specified collection into this list at the specified position
	AddAll(index int,  c Collection) bool
}

// AbstractList this class provides a skeletal implementation of the List interface to minimize the effort required
// to implement this interface
type AbstractList struct {

}






