package stack

import "testing"

func TestStackCreate(t *testing.T) {
	m := 3
	s := NewStack(m)
	if s.Top!=0 {
		t.Errorf("new stack top is not 0")
	}

	if s.Max!=m {
		t.Errorf("new stack max is not %d, but got %d", m, s.Max)
	}

	if len(s.Items)!=m {
		t.Errorf("new stack items array length should be %d, but got %d", m, len(s.Items))
	}
}

func TestStackTop(t *testing.T) {
	m := 5
	s := NewStack(m)

	s.Push(2)
	s.Push(4)
	s.Push(8)
	
	want := 3
	if s.Top!=want {
		t.Errorf("expected top %d, got %d", want, s.Top)
	}
}

func TestStackPeek(t *testing.T) {
	m := 5
	s := NewStack(m)
	s.Push(2)
	s.Push(4)
	s.Push(8)
	
	want := 8
	if got:=s.Items[s.Top-1]; got!=want {
		t.Errorf("expected top item %d, got %d", want, got)
	}
}

func TestStackOverflow(t *testing.T) {
	m := 5
	s := NewStack(m)

	s.Push(2)
	s.Push(4)
	s.Push(8)
	s.Push(10)
	s.Push(12)
	ok := s.Push(20)

	want := false
	if ok != want {
		t.Errorf("stack overflow expected with s.Push()=%t, got %t", want, ok)
	}
}

func TestStackPop(t *testing.T) {
	m := 5
	s := NewStack(m)
	s.Push(2)
	s.Push(4)
	s.Push(8)

	item, _ := s.Pop()

	want := 8
	if got:=item; got!=want {
		t.Errorf("expected pop item %d, got %d", want, got)
	}
}

func TestStackUnderflow(t *testing.T) {
	m := 5
	s := NewStack(m)
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
