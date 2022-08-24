package utils

type Element struct {
	// The pointer next is the linked to next element of the linked list single
	next *Element
	// The list in which this element belongs
	list *List
	// The value that this object stores
	Value interface {
	}
}

func (e *Element) Next() *Element {
	return e.next
}
