package queue

import "errors"

var (
	ErrOverflow     = errors.New("queue overflow")
	EOQ             = errors.New("end of queue")
	ErrNegativeSize = errors.New("negative queue size")
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

func New[V any](size int64) (*queue[V], error) {
	switch {
	case size >= -1:
		return &queue[V]{length: 0, size: size, front: nil, back: nil}, nil
	default:
		return nil, ErrNegativeSize
	}
}

func (q *queue[V]) Enqueue(v V) error {
	if q.length == q.size {
		return ErrOverflow
	}

	if q.front == nil {
		q.front = &node[V]{value: v, next: nil}
		q.back = q.front
		q.length++
		return nil
	}

	q.back.next = &node[V]{value: v, next: nil}
	q.back = q.back.next
	q.length++
	return nil
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

func (q *queue[V]) Peek() (V, error) {
	if q.front == nil {
		return *new(V), EOQ
	}
	return q.front.value, nil
}
