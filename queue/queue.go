package queue

type node[V any] struct {
	value V
	next  *node[V]
}

type queue[V any] struct {
	front, back *node[V]
}

func New[V any]() *queue[V] {
	return &queue[V]{front: nil, back: nil}
}

func (q *queue[V]) Enqueue(v V) {
	if q.front == nil {
		q.front = &node[V]{value: v, next: nil}
		q.back = q.front
		return
	}

	q.back.next = &node[V]{value: v, next: nil}
	q.back = q.back.next
}

func (q *queue[V]) Dequeue() V {
	if q.front == nil {
		return *new(V)
	}

	value := q.front.value
	q.front = q.front.next

	if q.front == nil {
		q.back = nil
	}

	return value
}

func (q *queue[V]) IsEmpty() bool {
	return q.front == nil
}

func (q *queue[V]) Peek() V {
	return q.front.value
}
