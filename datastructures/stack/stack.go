package stack

type Stack struct {
	top    *Element
	length int
}

type Element struct {
	value interface{}
	prev  *Element
}

// Create new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Get length of stack
func (s *Stack) Len() int {
	return s.length
}

// Get top element of stack
func (s *Stack) Top() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// Push a element to top of stack
func (s *Stack) Push(v interface{}) {
	s.top = &Element{v, s.top}
	s.length++
}

// Pop a element of top of stack
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	top := s.top.value
	s.top = s.top.prev
	s.length--
	return top
}
