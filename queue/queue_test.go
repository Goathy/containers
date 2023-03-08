package queue_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Goathy/containers/queue"
)

func TestQueue(t *testing.T) {
	t.Run("should create stack instance", func(t *testing.T) {
		q := queue.New[any]()

		assertBool(t, q == nil, "queue should not be nil")
	})

	t.Run("should insert items into queue", func(t *testing.T) {
		q := queue.New[int]()

		input := []int{1, 2, 3, 4, 5}

		for _, in := range input {
			q.Enqueue(in)
		}
	})

	t.Run("should remove items from queue", func(t *testing.T) {
		q := queue.New[string]()

		input := []string{"a", "b", "c", "d", "e"}

		for _, in := range input {
			q.Enqueue(in)
		}

		got := make([]string, 0)

		for !q.IsEmpty() {
			v := q.Dequeue()

			got = append(got, v)
		}

		assertBool(t, !reflect.DeepEqual(got, input), fmt.Sprintf("want %q, got %q", input, got))
	})

	t.Run("should always peek first element from queue", func(t *testing.T) {

		tt := []struct {
			desc  string
			input []int
			want  int
		}{
			{
				desc:  "should pick 1",
				input: []int{1, 2, 3, 4, 5},
				want:  1,
			},
			{
				desc:  "should pick 10",
				input: []int{10},
				want:  10,
			},
		}

		for _, tc := range tt {
			t.Run(tc.desc, func(t *testing.T) {
				q := queue.New[int]()

				for _, in := range tc.input {
					q.Enqueue(in)
				}

				got := q.Peek()

				assertValue(t, got, tc.want)
			})
		}

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
			q := queue.New[string]()

			for _, in := range tc.input {
				q.Enqueue(in)
			}

			assertBool(t, q.IsEmpty() != tc.want, fmt.Sprintf("want %t, got %t", tc.want, q.IsEmpty()))
		})
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
