package main

// Stack is a limited FIFO for holding sparkline values.
type Stack struct {
	Values []int
	Len    int
}

// NewStack inits new Stack with size limit.
func NewStack(size int) *Stack {
	return &Stack{
		Values: make([]int, size),
		Len:    size,
	}
}

// Push inserts data to stack, preserving constant length.
func (s *Stack) Push(val int) {
	s.Values = append(s.Values, val)
	if len(s.Values) > s.Len {
		s.Values = s.Values[1:]
	}
}

// Front returns front value.
func (s *Stack) Front() int {
	if len(s.Values) == 0 {
		return 0
	}
	return s.Values[len(s.Values)-1]
}