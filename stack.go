package stack

import (
	"errors"
)

var (
	ErrStackOverflow = errors.New("stack overflow")
	// EOS error is returned when Pop() or Peek() method are called
	// on empty stack
	EOS                  = errors.New("EOS")
	ErrNegativeStackSize = errors.New("negative stack size")
)

type node[Value any] struct {
	value Value
	prev  *node[Value]
}

type stack[Value any] struct {
	size   int64
	length int64
	top    *node[Value]
}

func New[Value any](size int64) (*stack[Value], error) {
	switch {
	case size == -1:
		return &stack[Value]{-1, 0, nil}, nil
	case size >= 0:
		return &stack[Value]{size, 0, nil}, nil
	default:
		return nil, ErrNegativeStackSize
	}
}

func (s *stack[Value]) IsFull() bool {
	if s.size == -1 {
		return false
	}
	return s.length == s.size
}

func (s *stack[Value]) IsEmpty() bool {
	return s.top == nil
}

func (s *stack[Value]) Push(element Value) error {
	if s.IsFull() {
		return ErrStackOverflow
	}

	if s.IsEmpty() {
		s.top = &node[Value]{element, nil}
		s.length++
		return nil
	}

	s.top = &node[Value]{element, s.top}
	s.length++

	return nil
}

func (s *stack[Value]) Pop() (Value, error) {
	if s.IsEmpty() {
		// *new(Value) create zero value of provided type, equivalent to e.g. var zero Value
		return *new(Value), EOS
	}

	n := s.top
	s.top = n.prev
	s.length--

	return n.value, nil
}

func (s *stack[Value]) Peek() (Value, error) {
	if s.IsEmpty() {
		return *new(Value), EOS
	}

	return s.top.value, nil
}
