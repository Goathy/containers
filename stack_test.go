package stack_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Goathy/stack"
)

func TestStack(t *testing.T) {
	t.Run("should create stack instance", func(t *testing.T) {
		s, err := stack.New[int64](0)

		assertBool(t, err != nil, "unexpected error")

		assertBool(t, s == nil, "stack should not be nil")
	})

	t.Run("should create stack without size limit", func(t *testing.T) {
		s, err := stack.New[float64](-1)

		assertBool(t, err != nil, "unexpected error")

		assertBool(t, !s.IsEmpty(), "empty stack should be empty")

		for _, el := range []float64{1, 0.2, 00.3, 4.01, 5.4, 345.6, 7} {
			s.Push(el)
		}

		assertBool(t, s.IsFull(), "stack with no size limit should not be full")
	})

	t.Run("should not create stack with size less then -1", func(t *testing.T) {
		s, err := stack.New[string](-2)

		want := errors.New("negative stack size")

		assertError(t, want, err)

		assertBool(t, s != nil, "stack should be nil")
	})

	t.Run("stack should be full", func(t *testing.T) {
		s, err := stack.New[rune](5)

		assertBool(t, err != nil, "unexpected error")

		assertBool(t, !s.IsEmpty(), "stack should be empty")

		for _, el := range []rune{'a', 'b', 'c', 'd', 'e'} {
			s.Push(el)
		}

		assertBool(t, !s.IsFull(), "stack should be full")
	})

	t.Run("should return error if provide more values then stack limit", func(t *testing.T) {
		s, err := stack.New[int](5)

		assertBool(t, err != nil, "unexpected error")

		for _, el := range []int{1, 2, 3, 4, 5} {
			err := s.Push(el)

			assertBool(t, err != nil, "unexpected error")
		}

		err = s.Push(6)

		want := errors.New("stack overflow")

		assertError(t, want, err)
	})

	t.Run("should return error if pop more values then stack store", func(t *testing.T) {
		s, err := stack.New[complex128](0)

		assertBool(t, err != nil, "unexpected error")

		_, err = s.Pop()

		assertError(t, stack.EOS, err)
	})

	t.Run("stack should pop values in revers order to push", func(t *testing.T) {
		type TestStruct struct {
			description string
			value       int
		}

		elements := []TestStruct{{"foo", 10}, {"bar", 20}, {"baz", 30}}

		s, err := stack.New[TestStruct](3)

		assertBool(t, err != nil, "unexpected error")

		for _, el := range elements {
			s.Push(el)
		}

		assertBool(t, !s.IsFull(), "stack should be full")

		for i := len(elements) - 1; i >= 0; i-- {
			el, err := s.Pop()

			assertBool(t, err != nil, "unexpected error")

			assertValue(t, el, elements[i])
		}
	})

	t.Run("peek should display value without popping value from stack", func(t *testing.T) {
		s, err := stack.New[float32](1)

		assertBool(t, err != nil, "unexpected error")

		s.Push(6.32)

		assertBool(t, !s.IsFull(), "stack should be full")

		got, err := s.Peek()

		want := float32(6.32)

		assertValue(t, got, want)

		assertBool(t, err != nil, "unexpected error")

		assertBool(t, !s.IsFull(), "stack should be full")

	})

	t.Run("peek on empty stack should return error", func(t *testing.T) {
		s, err := stack.New[float32](0)

		assertBool(t, err != nil, "unexpected error")

		assertBool(t, !s.IsFull(), "stack should be full")

		_, err = s.Peek()

		assertError(t, stack.EOS, err)

		assertBool(t, !s.IsFull(), "stack should be full")
	})
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
