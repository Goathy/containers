package queue_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Goathy/containers/queue"
)

func TestQueue(t *testing.T) {
	t.Run("should create stack instance", func(t *testing.T) {
		q, err := queue.New[any](0)

		assertBool(t, err != nil, "unexpected error")

		assertBool(t, q == nil, "queue should not be nil")
	})

	t.Run("should return error if size is less then -1", func(t *testing.T) {
		q, err := queue.New[int](-2)

		assertError(t, queue.ErrNegativeSize, err)

		assertBool(t, q != nil, "queue should be nil")

	})

	t.Run("should create queue without size when -1 is provided", func(t *testing.T) {
		q, err := queue.New[float64](-1)

		assertBool(t, err != nil, "unexpected error")

		input := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		for _, in := range input {
			q.Enqueue(in)
		}

		want := false

		assertValue(t, q.IsFull(), want)
	})

	t.Run("should return error if enqueue beyond queue size", func(t *testing.T) {
		q, err := queue.New[uint16](9)

		assertBool(t, err != nil, "unexpected error")

		input := []uint16{1, 2, 3, 4, 5, 6, 7, 8, 9}

		for _, in := range input {
			q.Enqueue(in)
		}

		err = q.Enqueue(10)

		assertError(t, queue.ErrOverflow, err)

		want := true

		assertValue(t, q.IsFull(), want)
	})

	t.Run("should insert items into queue", func(t *testing.T) {
		q, err := queue.New[int](5)

		assertBool(t, err != nil, "unexpected error")

		input := []int{1, 2, 3, 4, 5}

		for _, in := range input {
			q.Enqueue(in)
		}
	})

	t.Run("should remove items from queue", func(t *testing.T) {
		q, err := queue.New[string](5)

		assertBool(t, err != nil, "unexpected error")

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
		q, err := queue.New[float32](10)

		assertBool(t, err != nil, "unexpected error")

		_, err = q.Dequeue()

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
			q, err := queue.New[string](3)

			assertBool(t, err != nil, "unexpected error")

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
			q, err := queue.New[rune](tc.size)

			assertBool(t, err != nil, "unexpected error")

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

func assertValue(t testing.TB, got any, want any) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %q, got %q", want, got)
	}
}
