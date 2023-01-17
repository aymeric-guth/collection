package stack

type Stack[T any] struct {
	d []T
}

func New[T any](args ...T) *Stack[T] {
	return &Stack[T]{d: args}
}

func (s *Stack[T]) Push(v T) {
	s.d = append(s.d, v)
}

func (s *Stack[T]) Pop() T {
	v := s.d[len(s.d)-1]
	s.d = s.d[:len(s.d)-1]
	return v
}

func (s *Stack[T]) Peek() T {
	return s.d[len(s.d)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.d)
}

func (s *Stack[T]) Empty() {
	panic("not implemented")
}
