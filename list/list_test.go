package list_test

import (
	"reflect"
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

func TestSearchOnEmpty(t *testing.T) {
	l := list.New[int]()

	value := l.Search(10)

	assertValue(t, value, 0)
}

func TestSearch(t *testing.T) {
	l := list.New[string]()

	input := []string{"a", "b", "c", "d", "e"}

	for _, in := range input {
		l.Insert(in)
	}

	got := l.Search("c")
	want := "c"

	assertValue(t, got, want)
}

func TestDeleteOnEmpty(t *testing.T) {
	l := list.New[int]()

	l.Delete(10)
}

func TestDelete(t *testing.T) {
	l := list.New[string]()

	input := []string{"a", "b", "c", "d", "e"}

	for _, in := range input {
		l.Insert(in)
	}

	l.Delete("c")

	got := l.Search("c")
	want := ""

	assertValue(t, got, want)
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
