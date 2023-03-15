package list

import "reflect"

type node[V any] struct {
	value V
	next  *node[V]
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
		l.head = &node[V]{value: v, next: nil}
		l.tail = l.head
		return
	}

	l.tail.next = &node[V]{value: v, next: nil}
	l.tail = l.tail.next
}

func (l *list[V]) Search(v V) V {
	if l.head == nil {
		var value V
		return value
	}

	n := l.head

	for n != nil && !reflect.DeepEqual(n.value, v) {
		n = n.next
	}

	if n == nil {
		var value V
		return value
	}

	return n.value
}

func (l *list[V]) Delete(v V) {
	if l.head == nil {
		return
	}

	n := l.head

	for n.next != nil && !reflect.DeepEqual(n.next.value, v) {
		n = n.next
	}
	n.next = n.next.next
}
