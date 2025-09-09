package stack

import (
	"sync"
)

type Element struct {
	Data any
	next *Element
}

type Stack struct {
	lock *sync.Mutex
	top  *Element
	Size int
}

func (s *Stack) Push(data any) {
	s.lock.Lock()
	defer s.lock.Unlock()

	elm := &Element{
		Data: data,
		next: nil,
	}

	if s.top == nil {
		s.top = elm
		s.Size++
		return
	}

	elm.next = s.top
	s.top = elm
	s.Size++

}

func (s *Stack) Pop() (any, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.top == nil {
		return nil, false
	}

	elm := s.top.Data

	temp := s.top
	s.top = s.top.next
	temp.next = nil
	s.Size--

	return elm, true
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func NewStack() *Stack {
	return &Stack{
		lock: &sync.Mutex{},
	}
}
