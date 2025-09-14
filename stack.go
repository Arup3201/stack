package stack

import "sync"

type itemType any

type stack struct {
	items  []itemType
	rwLock sync.RWMutex
}

func MakeStack() *stack {
	return &stack{
		items:  make([]itemType, 0),
		rwLock: sync.RWMutex{},
	}
}

func (s *stack) Push(item itemType) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.items = append(s.items, item)
}

func (s *stack) Pop() (*itemType, bool) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	item, ok := s.Peek()
	if !ok {
		return nil, false
	}
	s.items = s.items[0 : len(s.items)-1]
	return item, true
}

func (s *stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *stack) Peek() (*itemType, bool) {
	if s.isEmpty() {
		return nil, false
	}

	return &s.items[len(s.items)-1], true
}
