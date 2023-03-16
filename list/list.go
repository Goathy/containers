package list

import "reflect"

type node[V any] struct {
	Value V
	next  *node[V]
}

func (n *node[V]) Next() *node[V] {
	if p := n.next; p != nil {
		return p
	}
	return nil
}

type list[V any] struct {
	head, tail *node[V]
}

func New[V any]() *list[V] {
	return &list[V]{head: nil, tail: nil}
}

func (l *list[V]) Insert(v V) {
	if l.head == nil {
		l.head = &node[V]{Value: v, next: nil}
		l.tail = l.head
		return
	}

	l.tail.next = &node[V]{Value: v, next: nil}
	l.tail = l.tail.next
}

func (l *list[V]) Search(v V) V {
	var zeroValue V

	if l.head == nil {
		return zeroValue
	}

	for n := l.Traverse(); n != nil; n = n.Next() {
		if reflect.DeepEqual(n.Value, v) {
			return n.Value
		}
	}

	return zeroValue
}

func (l *list[V]) Delete(v V) {
	defer func() {
		if l.head == nil {
			l.tail = nil
		}
	}()

	if l.head == nil {
		return
	}

	if reflect.DeepEqual(l.head.Value, v) {
		l.head = l.head.next
		return
	}

	var n *node[V]

	for n = l.Traverse(); n.next != nil && !reflect.DeepEqual(n.next.Value, v); n = n.Next() {
	}

	if n.next == nil {
		return
	}

	n.next = n.next.next
}

func (l *list[V]) Traverse() *node[V] {
	if l.head == nil {
		return nil
	}

	return l.head
}
