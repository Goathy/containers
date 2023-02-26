package stack_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Goathy/stack"
)

func TestStack(t *testing.T) {
	t.Run("should create stack instance", func(t *testing.T) {
		stack, err := stack.New[int64](0)

		assertBool(t, err != nil, "unexpected error")

		assertBool(t, stack == nil, "stack should not be nil")
	})

	t.Run("should create stack without size limit", func(t *testing.T) {
		stack, err := stack.New[int64](-1)

		assertBool(t, err != nil, "unexpected error")

		assertBool(t, !stack.IsEmpty(), "empty stack should be empty")

		for _, el := range []int64{1, 2, 3, 4, 5, 6, 7} {
			stack.Push(el)
		}

		assertBool(t, stack.IsFull(), "stack with no size limit should not be full")
	})

	t.Run("should not create stack with size less then -1", func(t *testing.T) {
		stack, err := stack.New[int64](-2)

		want := errors.New("negative stack size")

		assertError(t, want, err)

		assertBool(t, stack != nil, "stack should be nil")
	})

	t.Run("stack should be full", func(t *testing.T) {
		stack, err := stack.New[int64](5)

		assertBool(t, err != nil, "unexpected error")

		assertBool(t, !stack.IsEmpty(), "stack should be empty")

		for _, el := range []int64{1, 2, 3, 4, 5} {
			stack.Push(el)
		}

		assertBool(t, !stack.IsFull(), "stack should be full")
	})

	t.Run("should return error if provide more values then stack limit", func(t *testing.T) {
		stack, err := stack.New[int64](5)

		assertBool(t, err != nil, "unexpected error")

		for _, el := range []int64{1, 2, 3, 4, 5} {
			err := stack.Push(el)

			assertBool(t, err != nil, "unexpected error")
		}

		err = stack.Push(6)

		want := errors.New("stack overflow")

		assertError(t, want, err)
	})

	t.Run("should return error if pop more values then stack store", func(t *testing.T) {
		stack, err := stack.New[int64](0)

		assertBool(t, err != nil, "unexpected error")

		_, err = stack.Pop()

		want := errors.New("stack empty")

		assertError(t, want, err)
	})

	t.Run("stack should pop values in reverso order to push", func(t *testing.T) {
		stack, err := stack.New[int64](5)

		assertBool(t, err != nil, "unexpected error")

		elements := []int64{1, 2, 3, 4, 5}

		for _, el := range elements {
			err := stack.Push(el)

			assertBool(t, err != nil, "unexpected error")
		}

		assertBool(t, !stack.IsFull(), "stack should be full")

		for i := len(elements) - 1; i >= 0; i-- {
			el, err := stack.Pop()

			assertBool(t, err != nil, "unexpected error")

			assertValues(t, el, elements[i])
		}

		assertBool(t, !stack.IsEmpty(), "stackshould be empty")
	})
}

func TestStackWithStrings(t *testing.T) {
	elements := []string{"foo", "bar", "baz"}

	stack, err := stack.New[string](3)

	assertBool(t, err != nil, "unexpected error")

	for _, el := range elements {
		stack.Push(el)
	}

	assertBool(t, !stack.IsFull(), "stack should be full")

	for i := len(elements) - 1; i >= 0; i-- {
		el, err := stack.Pop()

		assertBool(t, err != nil, "unexpected error")

		assertValues(t, el, elements[i])
	}
}

func TestStackWithStruct(t *testing.T) {
	type TestStruct struct {
		description string
		value       int
	}

	elements := []TestStruct{{"foo", 10}, {"bar", 20}, {"baz", 30}}

	stack, err := stack.New[TestStruct](3)

	assertBool(t, err != nil, "unexpected error")

	for _, el := range elements {
		stack.Push(el)
	}

	assertBool(t, !stack.IsFull(), "stack should be full")

	for i := len(elements) - 1; i >= 0; i-- {
		el, err := stack.Pop()

		assertBool(t, err != nil, "unexpected error")

		assertValues(t, el, elements[i])
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

func assertValues(t testing.TB, got any, want any) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %q, got %q", want, got)
	}
}
