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
	head *node[V]
	tail *node[V]
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
	if l.head == nil {
		var value V
		return value
	}

	n := l.head

	for n != nil && !reflect.DeepEqual(n.Value, v) {
		n = n.next
	}

	if n == nil {
		var value V
		return value
	}

	return n.Value
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

	n := l.head

	if reflect.DeepEqual(n.Value, v) {
		l.head = n.next

		return
	}

	for n.next != nil && !reflect.DeepEqual(n.next.Value, v) {
		n = n.next
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
