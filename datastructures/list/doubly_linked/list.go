package doubly_linked

type List struct {
	head, tail *Node
	length     int
}

// Node of doubly linked list
type Node struct {
	value      interface{}
	next, prev *Node
}

// Get length of list
func (l *List) Len() int {
	return l.length
}

// Get the first element of list
func (l *List) Front() interface{} {
	if l.length == 0 {
		return nil
	}
	return l.head.value
}

// Get the last element of list
func (l *List) Back() interface{} {
	if l.length == 0 {
		return nil
	}
	return l.tail.value
}

func (l *List) GetByIdx(idx int) interface{} {
	if idx < 0 || idx >= l.length {
		return nil
	}
	// 	cur := l.head
	// 	for i := 0; i < idx; i++ {
	// 		cur = cur.next
	// 	}
	// 	return cur.value
	return nil
}

func (l *List) IsExist(val interface{}) bool {
	return false
}
