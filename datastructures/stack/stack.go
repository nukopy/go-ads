package stack

type Stack struct {
	top    *Element
	length int
}

// Element of stack
type Element struct {
	value interface{}
	prev  *Element
}

// Create new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Get length of stack
// Complexity: O(1)
func (s *Stack) Len() int {
	return s.length
}

// Check if stack is empty
// Complexity: O(1)
func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

// Get top element of stack
// Complexity: O(1)
func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.top.value
}

// Push a element to top of stack
// Complexity: O(1)
func (s *Stack) Push(v interface{}) {
	s.top = &Element{v, s.top}
	s.length++
}

// Pop a element from stack
// Complexity: O(1)
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	top := s.top.value
	s.top = s.top.prev
	s.length--
	return top
}
