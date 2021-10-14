package datastructures

type Stack struct {}

func (s *Stack) New() *Stack {
	return &Stack{}
}

func (s *Stack) Len() int {
	// return length of stack
	return 0
}

func (s *Stack) Push(v interface{}) {
	// push a element to top of stack
}

func (s *Stack) Pop() interface{} {
	// pop a element of top of stack
	return nil
}
