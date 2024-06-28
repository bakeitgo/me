package queue

type node[T any] struct {
	data T
	next *node[T]
}

type queue[T any] struct {
	tail *node[T]
	head *node[T]
	len  int
}

func (q *queue[T]) Peek() (firstItem T) {
	firstItem = q.head.data
	return
}

func (q *queue[T]) Deque() (dequed T) {
	dequed, q.head = q.head.data, q.head.next
	q.len--
	return
}

func (q *queue[T]) Enqueue(item T) {
	node := &node[T]{item, nil}
	if q.len == 0 {
		q.tail = node
		q.head = node
		return
	}
	q.tail.next = node
	q.tail = node
	defer func() { q.len++ }()
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{}
}
