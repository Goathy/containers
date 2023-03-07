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

		for {
			v := q.Dequeue()

			if v == "" {
				break
			}

			got = append(got, v)
		}

		assertBool(t, !reflect.DeepEqual(got, input), fmt.Sprintf("want %q, got %q", input, got))
	})
}

func assertBool(t testing.TB, got bool, msg string) {
	t.Helper()
	if got {
		t.Error(msg)
	}
}
