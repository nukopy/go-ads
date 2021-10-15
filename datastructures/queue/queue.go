package queue

import "fmt"

type Queue struct {
	front  *Element
	back   *Element
	length int
}

// Element of queue
type Element struct {
	value interface{}
	next  *Element
}

// Create new queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// Check if queue is empty
// Complexity: O(1)
func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

// Get length of queue
// Complexity: O(1)
func (q *Queue) Len() int {
	return q.length
}

// Push a element to back of queue
// Complexity: O(1)
func (q *Queue) Enqueue(v interface{}) {
	e := &Element{v, nil}

	if q.IsEmpty() {
		q.front = e
	} else if q.length == 1 {
		q.front.next = e
		q.back = e
	} else {
		// length > 1
		q.back.next = e // (&q).(&back).next = e
		q.back = e
		// 一見、q.back が上書きされて、どこからもアクセスできなくなり消えてしまうように見えるけど、
		// back のポインタ経由で q.back.next を書き換えてるから、q.back を更新しても書き換えた結果は残る。
		// 間に挟まった要素（今回消えてしまったように見えた要素）は front から辿ればアクセスできるようになっている。
	}
	q.length++
}

// Extract a element from Queue
// Complexity: O(1)
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	// update q.front
	f := q.front.value
	q.front = q.front.next
	q.length--
	return f
}

// Get next element in queue
// Complexity: O(1)
func (q *Queue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.front.value
}

// Get last element of Queue
// Complexity: O(1)
func (q *Queue) Back() interface{} {
	if q.IsEmpty() {
		return nil
	} else if q.length == 1 {
		return q.front.value
	}
	return q.back.value
}

func (q *Queue) printAll() {
	for e := q.front; e != nil; e = e.next {
		fmt.Printf("%v\n", e)
	}
	for !q.IsEmpty() {
		fmt.Printf("%d ", q.Dequeue())
	}
}
