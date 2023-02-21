package util

type Stack[T any] []T

func CreateStack[T any](l int) *Stack[T] {
	s := make(Stack[T], 0, l)
	return &s
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func (s *Stack[T]) IsEmpty() bool {
	return (*s).Size() == 0
}

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() (v T, ok bool) {
	if s.IsEmpty() {
		return
	}
	i := s.Size() - 1
	v = (*s)[i]
	*s = (*s)[:i]
	return v, true
}

func (s *Stack[T]) GetLast() (v T, ok bool) {
	if s.IsEmpty() {
		return
	}
	i := s.Size() - 1
	v = (*s)[i]
	return v, true
}
