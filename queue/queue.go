package queue

type Queue[T any] struct {
	q []T
}

type IQueue[T any] interface {
	Push(Queue[T])
	Pop() Queue[T]
	Peek() Queue[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{q: make([]T, 0)}
}

func (q *Queue[T]) Push(v T) {
	q.q = append(q.q, v)
}

func (q *Queue[T]) Pop() *T {
	if len(q.q) > 0 {
		v := q.q[0]
		q.q = q.q[1:]
		return &v
	}
	return nil
}

func (q *Queue[T]) Peek() *T {
	if len(q.q) > 0 {
		return &q.q[0]
	}
	return nil
}