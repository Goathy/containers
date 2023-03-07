package queue

import "errors"

var (
	EOQ = errors.New("end of queue")
)

type node[V any] struct {
	value V
	next  *node[V]
}

type queue[V any] struct {
	front  *node[V]
	back   *node[V]
	length int64
	size   int64
}

func New[V any](size int64) *queue[V] {
	return &queue[V]{length: 0, size: size}
}

func (q *queue[V]) Enqueue(v V) {
	defer func() { q.length++ }()
	if q.front == nil {
		q.front = &node[V]{value: v, next: nil}
		q.back = q.front
		return
	}

	q.back.next = &node[V]{value: v, next: nil}
	q.back = q.back.next
}

func (q *queue[V]) Dequeue() (V, error) {
	defer func() {
		if q.front == nil {
			q.back = q.front
		}
	}()

	if q.front == nil {
		return *new(V), EOQ
	}

	value := q.front.value
	q.front = q.front.next

	return value, nil
}

func (q *queue[V]) IsEmpty() bool {
	return q.front == nil
}

func (q *queue[V]) IsFull() bool {
	return q.length == q.size
}
