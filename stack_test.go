package stack

import "testing"

func assert(t testing.TB, msg string, got, want interface{}) {
	t.Helper()

	if got == want {
		return
	}

	var errStr string
	switch got.(type) {
	case int:
		errStr = msg + ", expected %d, got %d"
		t.Errorf(errStr, got, want)
	case string:
		errStr = msg + ", expected %s, got %s"
		t.Errorf(errStr, got, want)
	default:
		errStr = msg + ", expected %v, got %v"
		t.Errorf(errStr, got, want)
	}
}

func TestStackPush(t *testing.T) {
	t.Run("push 10", func(t *testing.T) {
		s := MakeStack()

		s.Push(10)

		want := 10
		got, _ := s.Peek()
		assert(t, "stack push failed", *got, want)
	})
	t.Run("push 10 20", func(t *testing.T) {
		s := MakeStack()

		s.Push(10)
		s.Push(20)

		want := 20
		got, _ := s.Peek()
		assert(t, "stack push failed", *got, want)
	})
}

func TestStackPop(t *testing.T) {
	t.Run("pop 30 from stack 10 20 30", func(t *testing.T) {
		s := MakeStack()
		s.Push(10)
		s.Push(20)
		s.Push(30)

		got, _ := s.Pop()

		want := 30
		assert(t, "stack pop failed", *got, want)
	})
}
