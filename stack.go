package stack

type Stack struct {
	Max int
	Top int
	Items []int
}

func NewStack(m int) *Stack {
	return &Stack{
		Max: m, 
		Top: 0, 
		Items: make([]int, m),
	}
}

func (s *Stack) Push(item int) bool {
	if s.Top==s.Max {
		return false
	} else {
		s.Items[s.Top] = item
		s.Top += 1
		return true
	}
}

func (s *Stack) Pop() (int, bool) {
	if s.Top==0 {
		return 0, false
	} else {
		s.Top -= 1
		item := s.Items[s.Top]
		return item, true
	}
}
