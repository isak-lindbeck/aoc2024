package utils

type Queue[T interface{}] struct {
	current, last int
	data          []T
}

func NewQueue[T comparable](data []T) Queue[T] {
	return Queue[T]{0, 0, data}
}

func (q *Queue[T]) Push(t T) {
	q.data[q.last] = t
	q.last = (q.last + 1) % len(q.data)
	if q.current == q.last {
		// TODO Resize?
		panic("queue is full")
	}
}
func (q *Queue[T]) Pop() (T, bool) {
	if q.current == q.last {
		return q.data[0], false
	}
	t := q.data[q.current]
	q.current = (q.current + 1) % len(q.data)
	return t, true
}
