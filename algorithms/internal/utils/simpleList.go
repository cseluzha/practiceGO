package utils

// Represents a simply linked list
type List struct {
	root *Element
	len  int
}

func New() *List {
	return new(List).Init()
}

func (l *List) Init() *List {
	l.len = 0
	l.root = nil
	return l
}

func (l *List) Len() int {
	return l.len
}

// Insert an element at the final of the list
func (l *List) PushBack(v interface{}) *Element {
	// Instance of an element
	el := &Element{
		next:  nil,
		list:  l,
		Value: v,
	}

	if l.root == nil {
		l.root = el
	} else {
		var aux *Element
		for aux = l.root; aux.next != nil; aux = aux.Next() {
		}
		aux.next = el
	}
	l.len++
	return el
}

func (l *List) PushFront(v interface{}) *Element {
	el := &Element{
		next:  nil,
		list:  l,
		Value: v,
	}

	if l.root == nil {
		l.root = el
	} else {
		el.next = l.root
		l.root = el
	}
	l.len++
	return l.root
}

func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		if l.root == e {
			l.root = e.next
		} else {
			var current *Element
			for current = l.Front(); current.next != e && current.next != nil; current = current.next {
			}
			current.next = e.next
		}
		l.len--
	}
	return e.Value
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	var current *Element
	for current = l.root; current.next != nil; current = current.Next() {
	}
	return current
}
