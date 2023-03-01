package stack

import (
	"errors"
)

var (
	ErrorOverflow = errors.New("stack overflow")
	// EOS error is returned when Pop() or Peek() method are called
	// on empty stack
	EOS             = errors.New("EOS")
	ErrNegativeSize = errors.New("negative stack size")
)

type node[V any] struct {
	value V
	prev  *node[V]
}

type stack[V any] struct {
	size   int64
	length int64
	top    *node[V]
}

func New[V any](size int64) (*stack[V], error) {
	if size == -1 || size >= 0 {
		return &stack[V]{size, 0, nil}, nil
	}
	return nil, ErrNegativeSize
}

func (s *stack[V]) IsFull() bool {
	if s.size == -1 {
		return false
	}
	return s.length == s.size
}

func (s *stack[V]) IsEmpty() bool {
	return s.top == nil
}

func (s *stack[V]) Push(element V) error {
	if s.IsFull() {
		return ErrorOverflow
	}

	if s.IsEmpty() {
		s.top = &node[V]{element, nil}
		s.length++
		return nil
	}

	s.top = &node[V]{element, s.top}
	s.length++

	return nil
}

func (s *stack[V]) Pop() (V, error) {
	if s.IsEmpty() {
		// *new(V) create zero value of provided type, equivalent to e.g. var zero Value
		return *new(V), EOS
	}

	n := s.top
	s.top = n.prev
	s.length--

	return n.value, nil
}

func (s *stack[V]) Peek() (V, error) {
	if s.IsEmpty() {
		return *new(V), EOS
	}

	return s.top.value, nil
}
