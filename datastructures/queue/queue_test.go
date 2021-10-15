package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := New()
	q.Enqueue(5)

	// test Front, Back when length of queue is 1
	if q.Front() != q.Back() {
		t.Errorf("Expected 1, got %d", q.Front())
	}

	q.Enqueue(7)
	q.Enqueue(9)

	// test Len
	// q: 5 7 9
	if q.Len() != 3 {
		t.Errorf("Length of queue should be 3, but got %d", q.Len())
	}

	// test Dequeue
	v1 := q.Dequeue()
	// q: 7 9
	if v1 != 5 {
		t.Errorf("Dequeue should return 5, but got %d", v1)
	}
	if q.Front() != 7 {
		t.Errorf("Front should return 7, but got %d", q.Front())
	}
	if q.Len() != 2 {
		t.Errorf("Length of queue should be 2, but got %d", q.Len())
	}

	// test IsEmpty
	q.Dequeue()
	q.Dequeue()
	if q.IsEmpty() == false {
		t.Errorf("Queue should be empty, but got %d", q.Len())
	}

	// test Dequeue when queue is empty
	v2 := q.Dequeue()
	if v2 != nil {
		t.Errorf("Dequeue from empty queue should return nil, but got %d", v2)
	}
}

func TestQueueInternal(t *testing.T) {
	q := New()

	// test first Enqueue
	q.Enqueue(5)
	f := q.front
	if f == nil && f.next != nil {
		t.Errorf("Front should NOT be nil & Next should be nil , got %+v", f)
	}

	// test second Enqueue to check if q.front.next and q.back are equal when q.length is 2
	q.Enqueue(7)
	n := q.front.next
	b := q.back
	if n != b {
		t.Errorf("Pointer n(q.front.next) and b(q.back) should be equal when q.length is 2, got n: %p, b: %p", n, b)
	}

	// test third Enqueue to check if q.front.next.next and q.back are equal when q.length is 3
	q.Enqueue(9)
	nn := q.front.next.next
	b = q.back
	if nn != b {
		t.Errorf("Pointer nn(q.front.next.next) and b(q.back) should be equal when q.length is 3, got nn: %p, b: %p", nn, b)
	}
}

func TestPrintQueue(t *testing.T) {
	q := New()
	q.Enqueue(5)
	q.Enqueue(7)
	q.Enqueue(9)
	q.Enqueue(11)
	q.Enqueue(13)
	q.Dequeue()
	q.printAll() // 7 9 11 13
}
