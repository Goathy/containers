package list

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
