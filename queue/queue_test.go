package queue_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Goathy/containers/queue"
)

func TestQueue(t *testing.T) {
	t.Run("should create stack instance", func(t *testing.T) {
		q := queue.New[any](0)

		assertBool(t, q == nil, "queue should not be nil")
	})

	t.Run("should insert items into queue", func(t *testing.T) {
		q := queue.New[int](5)

		input := []int{1, 2, 3, 4, 5}

		for _, in := range input {
			q.Enqueue(in)
		}
	})

	t.Run("should remove items from queue", func(t *testing.T) {
		q := queue.New[string](5)

		input := []string{"a", "b", "c", "d", "e"}

		for _, in := range input {
			q.Enqueue(in)
		}

		got := make([]string, 0)

		for {
			v, err := q.Dequeue()

			if err == queue.EOQ {
				break
			}

			got = append(got, v)
		}

		assertBool(t, !reflect.DeepEqual(got, input), fmt.Sprintf("want %q, got %q", input, got))
	})

	t.Run("should return error if dequeue from empty queue", func(t *testing.T) {
		q := queue.New[float32](10)

		_, err := q.Dequeue()

		assertError(t, queue.EOQ, err)
	})

}

func TestIsEmpty(t *testing.T) {
	tt := []struct {
		desc  string
		input []string
		want  bool
	}{
		{
			desc:  "queue should be empty",
			input: []string{},
			want:  true,
		},
		{
			desc:  "queue should not be empty",
			input: []string{"a", "b"},
			want:  false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			q := queue.New[string](3)

			for _, in := range tc.input {
				q.Enqueue(in)
			}

			assertBool(t, q.IsEmpty() != tc.want, fmt.Sprintf("want %t, got %t", tc.want, q.IsEmpty()))
		})
	}
}

func TestIsFull(t *testing.T) {
	tt := []struct {
		desc  string
		input []rune
		size  int64
		want  bool
	}{
		{
			desc:  "queue should not be full",
			input: []rune{'1', '2', 'a', 'b'},
			size:  5,
			want:  false,
		},
		{
			desc:  "queue should be full",
			input: []rune{'1'},
			size:  1,
			want:  true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			q := queue.New[rune](tc.size)

			for _, in := range tc.input {
				q.Enqueue(in)
			}

			assertBool(t, q.IsFull() != tc.want, fmt.Sprintf("want %t, got %t", tc.want, q.IsFull()))
		})
	}
}

func assertError(t testing.TB, want, got error) {
	t.Helper()

	if want.Error() != got.Error() {
		t.Errorf("want %q, got %q", want, got)
	}
}

func assertBool(t testing.TB, got bool, msg string) {
	t.Helper()
	if got {
		t.Error(msg)
	}
}
