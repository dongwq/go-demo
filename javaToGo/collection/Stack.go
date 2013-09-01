package collection

type Stack struct {
	data []interface{}
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() interface{} {
	last := s.Size() - 1
	if last >= 0 {
		x := s.data[last]
		s.data[last] = nil
		s.data = s.data[:last]
		return x
	}
	return nil
}
