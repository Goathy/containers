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

type node struct {
	value int64
	next  *node
}

type stack struct {
	size        int64
	currentSize int64
	head        *node
}

func New(size int64) (*stack, error) {
	switch {
	case size == -1:
		return &stack{-1, 0, nil}, nil
	case size < 0:
		return nil, errStackSize
	default:
		return &stack{size, 0, nil}, nil
	}
}

func (s *stack) IsFull() bool {
	if s.size == -1 {
		return false
	}
	return s.currentSize == s.size
}

func (s *stack) IsEmpty() bool {
	return s.head == nil
}

func (s *stack) Push(element int64) error {
	if s.IsFull() {
		return errStackOverflow
	}

	if s.IsEmpty() {
		s.head = &node{element, nil}
		s.currentSize++
		return nil
	}

	s.head = &node{element, s.head}
	s.currentSize++

	return nil
}

func (s *stack) Pop() (int64, error) {
	if s.IsEmpty() {
		return 0, errStackEmpty
	}

	curr := s.head
	s.head = curr.next
	s.currentSize--

	return curr.value, nil
}
