package stack

import "testing"

func TestStackCreate(t *testing.T) {
	s := NewStack()
	if !s.IsEmpty() {
		t.Errorf("new stack top is not nil")
	}
}

func TestStackSize(t *testing.T) {
	s := NewStack()
	s.Push(2)
	s.Push(4)
	s.Push(8)

	want := 3
	if got := s.Size; got != want {
		t.Errorf("expected stack size %d, got %d", want, got)
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack()
	s.Push(2)
	s.Push(4)
	s.Push(8)

	item, _ := s.Pop()

	want := 8
	if got := item; got != want {
		t.Errorf("expected pop item %d, got %d", want, got)
	}
}

func TestStackUnderflow(t *testing.T) {
	s := NewStack()
	s.Push(2)
	s.Push(4)
	s.Push(8)
	s.Pop()
	s.Pop()
	s.Pop()

	_, ok := s.Pop()

	want := false
	if ok != want {
		t.Errorf("expected stack underflow with s.Pop()=%t, got %t", want, ok)
	}
}
