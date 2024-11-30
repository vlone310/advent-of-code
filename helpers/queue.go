package helpers

type Queue[T any] []T

func (q *Queue[T]) Push(item T) {
	*q = append(*q, item)
}
func (q *Queue[T]) Pop() T {
	h := *q
	var el T
	el, *q = h[0], h[1:]
	return el
}
