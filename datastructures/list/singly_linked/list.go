package singly_linked

type List struct {
	head, tail *Node
	length     int
}

// Node of singly linked list
type Node struct {
	value interface{}
	next  *Node
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

// Get the value of element at index
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

func (l *DoublyLinkedList) IsExist(val interface{}) bool {
	return false
}
