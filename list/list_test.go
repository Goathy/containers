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

func TestDeleteOutsideRangeFromNotEmptyList(t *testing.T) {
	l := list.New[string]()

	input := []string{"a", "b", "c", "d", "e"}

	for _, in := range input {
		l.Insert(in)
	}

	l.Delete("f")
}

func TestDeleteOnFirstElement(t *testing.T) {
	l := list.New[string]()

	input := []string{"a", "b", "c", "d", "e"}

	for _, in := range input {
		l.Insert(in)
	}

	l.Delete("a")

	got := l.Search("a")
	want := ""

	assertValue(t, got, want)
}

func TestDeleteAllElements(t *testing.T) {
	l := list.New[string]()

	input := []string{"a", "b", "c", "d", "e"}

	for _, in := range input {
		l.Insert(in)
	}

	for _, in := range input {
		l.Delete(in)
	}

	want := ""

	for _, in := range input {
		got := l.Search(in)
		assertValue(t, got, want)
	}

}

func TestDeleteAllElementsInReverseOrder(t *testing.T) {
	l := list.New[string]()

	input := []string{"a", "b", "c", "d", "e"}

	for _, in := range input {
		l.Insert(in)
	}

	for i := len(input) - 1; i >= 0; i-- {
		l.Delete(string(input[i]))
	}

	want := ""

	for _, in := range input {
		got := l.Search(in)
		assertValue(t, got, want)
	}
}

func TestTraverse(t *testing.T) {
	l := list.New[string]()

	input := []string{"a", "b", "c", "d", "e"}

	for _, in := range input {
		l.Insert(in)
	}

	var i int
	for n := l.Traverse(); n != nil; n = n.Next() {
		assertValue(t, n.Value, input[i])
		i++
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
