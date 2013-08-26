
package collection

import(
	"fmt"
)

type Stack interface{
	data [] interface{}
}

func (s* Stack) Push(v interface{}) bool{
	s.data = append(s.data, v)
	
}
func (s *Stack) Pop() v interface{}{
	i := len(s.data) - 1
	if i < 0 {
		return nil
	}
	
	res := s.data[i]
	s.data[i] = nil
	s.data = s.data[:i]
	return res
}