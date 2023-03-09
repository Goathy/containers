package stack_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Goathy/containers/stack"
)

func TestStackConstructor(t *testing.T) {
	s := stack.New[any]()

	assertBool(t, s == nil, "stack should not be nil")
}

func TestPush(t *testing.T) {
	s := stack.New[int]()

	input := []int{1, 2, 3, 4, 5}

	for _, in := range input {
		s.Push(in)
	}
}

func TestPopOnEmpty(t *testing.T) {
	s := stack.New[any]()

	value := s.Pop()

	assertValue(t, value, nil)
}

func TestPop(t *testing.T) {
	s := stack.New[string]()

	for _, v := range []string{"a", "b", "c", "d", "e"} {
		s.Push(v)
	}

	want := []string{"e", "d", "c", "b", "a"}
	got := make([]string, 0)

	for !s.IsEmpty() {
		v := s.Pop()

		got = append(got, v)
	}

	assertBool(t, !reflect.DeepEqual(got, want), fmt.Sprintf("want %q, got %q", want, got))
}

func TestPeekOnEmpty(t *testing.T) {
	s := stack.New[any]()

	value := s.Peek()

	assertValue(t, value, nil)
}

func TestPeek(t *testing.T) {
	tt := []struct {
		desc  string
		input []int
		want  int
	}{
		{
			desc:  "should pick 5",
			input: []int{1, 2, 3, 4, 5},
			want:  5,
		},
		{
			desc:  "should pick 10",
			input: []int{10},
			want:  10,
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			s := stack.New[int]()

			for _, in := range tc.input {
				s.Push(in)
			}

			got := s.Peek()

			assertValue(t, got, tc.want)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tt := []struct {
		desc  string
		input []string
		want  bool
	}{
		{
			desc:  "stack should be empty",
			input: []string{},
			want:  true,
		},
		{
			desc:  "stack should not be empty",
			input: []string{"a", "b"},
			want:  false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			s := stack.New[string]()

			for _, in := range tc.input {
				s.Push(in)
			}

			assertBool(t, s.IsEmpty() != tc.want, fmt.Sprintf("want %t, got %t", tc.want, s.IsEmpty()))
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
