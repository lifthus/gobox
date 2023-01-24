package stackNonGeneric

type Stack struct {
	vals []interface{}
}

func (s *Stack) Push(val interface{}) {
	s.vals = append(s.vals, val)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.vals) == 0 {
		return nil, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}
