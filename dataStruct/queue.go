package datastruct

// Queue ...
type Queue struct {
	ll *LinkedList
}

// NewQueue ...
func NewQueue() *Queue {
	return &Queue{ll: &LinkedList{}}
}

// Push ...
func (q *Queue) Push(val int) {
	q.ll.AddNode(val)
}

// Pop ...
func (q *Queue) Pop() int {
	front := q.ll.Front()
	q.ll.PopFront()
	return front
}

// Empty ...
func (q *Queue) Empty() bool {
	return q.ll.Empty()
}