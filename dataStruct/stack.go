package datastruct

// Stack ...
type Stack struct {
	ll *LinkedList
}

// NewStack ...
func NewStack() *Stack {
	return &Stack{ll: &LinkedList{}}
}

// Empty ...
func (s *Stack) Empty() bool {
	return s.ll.Empty()
}

// Push ...
func (s *Stack) Push(val int) {
	s.ll.AddNode(val)
}

// Pop ...
func (s *Stack) Pop() int {
	back := s.ll.Back()
	s.ll.PopBack()
	return back
}