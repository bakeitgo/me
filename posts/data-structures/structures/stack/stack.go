package stack

import (
	"errors"
)

type node[T any] struct {
	data T
	prev *node[T]
}

type stack[T any] struct {
	head *node[T]
	len  int
}

func (s *stack[T]) Push(item T) {
	node := &node[T]{item, s.head}
	s.head = node
	s.len++
}

func (s *stack[T]) Pop() (T, error) {
	if s.len == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	item := s.head.data
	s.head = s.head.prev
	s.len--
	return item, nil
}

func (s *stack[T]) Peek() T {
	return s.head.data
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{}
}

/*
func main() {
	stack := NewStack[int]()

	stack.Push(1)
	fmt.Println(stack.Peek())
	stack.Push(2)
	fmt.Println(stack.Peek())

	if item, err := stack.Pop(); err != nil {
		fmt.Println(item)
	}
	if item, err := stack.Pop(); err != nil {
		fmt.Println(item)
	}
	item, err := stack.Pop()
	if err != nil {
		fmt.Println(item)
	}
	fmt.Println(fmt.Sprintf("error %s", err.Error()))
}
*/
