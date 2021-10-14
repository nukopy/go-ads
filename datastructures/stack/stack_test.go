package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := New()

	// test Push & Len
	s.Push(1)
	s.Push(3)
	s.Push(5)
	if s.Len() != 3 {
		t.Errorf("Stack length should be 3, got %d", s.length)
	}

	// test Top
	if s.Top() != 5 {
		t.Errorf("Stack top should be 5, got %d", s.Top())
	}

	// test Pop
	s.Pop()
	v := s.Pop()
	if v == nil || v != 3 {
		t.Errorf("Stack pop should be 3, got %d", v)
	}
	if s.Len() != 1 {
		t.Errorf("Stack length should be 1, got %d", s.Len())
	}
}
