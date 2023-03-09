package stack

type node[V any] struct {
	value V
	prev  *node[V]
}

type stack[V any] struct {
	top *node[V]
}

func New[V any]() *stack[V] {
	return &stack[V]{top: nil}
}

func (s *stack[V]) IsEmpty() bool {
	return s.top == nil
}

func (s *stack[V]) Push(v V) {
	if s.top == nil {
		s.top = &node[V]{value: v, prev: nil}
		return
	}

	s.top = &node[V]{value: v, prev: s.top}
}

func (s *stack[V]) Pop() V {
	if s.top == nil {
		return *new(V)
	}

	value := s.top.value
	s.top = s.top.prev

	return value
}

func (s *stack[V]) Peek() V {
	if s.top == nil {
		return *new(V)
	}

	return s.top.value
}
