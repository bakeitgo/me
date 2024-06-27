package queue

type node[T] struct {
	data T
	next *Node[T]
}

type queue[T] struct {
	tail *node[T]
	head *node[T]
	len int
}



func (q *queue) Peek() (firstItem T) {
	firstItem = s.head.data
	return
}

func (q *queue) Deque() (dequed T)  {
	dequed, q.head = q.head.data, q.head.next
	q.len--
	return
}

func (q *queue) Enqueue(item T) {
	node := &node[T]{item}
	if q.len == 0 {
		q.tail = node
		q.head = node
		return
	} 
	q.tail.next = node
	q.tail = node
	defer func() { q.len++ }()
}

func (q *queue) NewQueue() *queue {
	return &queue{}
}
