package list_test

import (
	"testing"

	"github.com/Goathy/containers/list"
)

func TestListNew(t *testing.T) {
	l := list.New[any]()

	assertBool(t, l == nil, "list should not be nil")
}

func assertBool(t testing.TB, got bool, msg string) {
	t.Helper()
	if got {
		t.Error(msg)
	}
}
