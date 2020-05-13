package stack

type Stack struct {
	list []int
}

func (s *Stack) Push(x int) {
	s.list = append(s.list, x)
}

func (s *Stack) Pop() int {
	result := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return result
}

func (s *Stack) Peek() int {
	return s.list[len(s.list)-1]
}

func (s *Stack) Empty() bool {
	return len(s.list) == 0
}
