package stack

import "errors"

const (
	stackOverflow     = "stack overflow"
	stackEmpty        = "stack empty"
	negativeStackSize = "negative stack size"
)

var (
	errStackOverflow = errors.New(stackOverflow)
	errStackEmpty    = errors.New(stackEmpty)
	errStackSize     = errors.New(negativeStackSize)
)

type node[Value any] struct {
	value Value
	next  *node[Value]
}

type stack[Value any] struct {
	size        int64
	currentSize int64
	head        *node[Value]
}

func New[Value any](size int64) (*stack[Value], error) {
	switch {
	case size == -1:
		return &stack[Value]{-1, 0, nil}, nil
	case size >= 0:
		return &stack[Value]{size, 0, nil}, nil
	default:
		return nil, errStackSize
	}
}

func (s *stack[Value]) IsFull() bool {
	if s.size == -1 {
		return false
	}
	return s.currentSize == s.size
}

func (s *stack[Value]) IsEmpty() bool {
	return s.head == nil
}

func (s *stack[Value]) Push(element Value) error {
	if s.IsFull() {
		return errStackOverflow
	}

	if s.IsEmpty() {
		s.head = &node[Value]{element, nil}
		s.currentSize++
		return nil
	}

	s.head = &node[Value]{element, s.head}
	s.currentSize++

	return nil
}

func (s *stack[Value]) Pop() (Value, error) {
	if s.IsEmpty() {
		// *new(Value) create zero value of provided type, equivalent to e.g. var zero Value
		return *new(Value), errStackEmpty
	}

	curr := s.head
	s.head = curr.next
	s.currentSize--

	return curr.value, nil
}
