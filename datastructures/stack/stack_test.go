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
		t.Errorf("Length of stack should be 3, got %d", s.length)
	}

	// test Top
	if s.Top() != 5 {
		t.Errorf("Top value of stack should be 5, got %d", s.Top())
	}

	// test Pop
	s.Pop()
	v1 := s.Pop()
	if v1 == nil || v1 != 3 {
		t.Errorf("v1 should be 3, got %d", v1)
	}
	if s.Len() != 1 {
		t.Errorf("Length of stack should be 1, got %d", s.Len())
	}

	// test Pop of empty stack
	s.Pop()
	v2 := s.Pop()
	if v2 != nil && s.Len() != 0 {
		t.Errorf("Value of pop of empty stack should be nil, got %d", v2)
	}
}
