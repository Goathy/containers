package list_test

import (
	"testing"

	"github.com/Goathy/containers/list"
)

func TestListNew(t *testing.T) {
	l := list.New[any]()

	assertBool(t, l == nil, "list should not be nil")
}

func TestInsert(t *testing.T) {
	l := list.New[int]()

	input := []int{1, 2, 3, 4, 5}

	for _, in := range input {
		l.Insert(in)
	}
}

func assertBool(t testing.TB, got bool, msg string) {
	t.Helper()
	if got {
		t.Error(msg)
	}
}
