package queue

type Queue[T any] struct {
	q []T
}

type IQueue[T any] interface {
	Push(Queue[T])
	Pop() Queue[T]
	Peek() Queue[T]
	Size() int
}

func New[T any](args ...T) *Queue[T] {
	return &Queue[T]{q: args}
}

func (q *Queue[T]) Push(v T) {
	q.q = append(q.q, v)
}

func (q *Queue[T]) Pop() T {
	v := q.q[0]
	q.q = q.q[1:]
	return v
}

func (q *Queue[T]) Peek() T {
	return q.q[0]
}

func (q *Queue[T]) Size() int {
	return len(q.q)
}
