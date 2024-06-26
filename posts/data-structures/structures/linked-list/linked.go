package linked_list

import (
	"errors"
)

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type SingleLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

func (s *SingleLinkedList[T]) InsertLast(data T) error {
	appendNode := &Node[T]{data, nil}
	defer func() { s.len += 1 }()
	if s.len == 0 {
		s.tail = appendNode
		s.head = appendNode
	} else {
		s.tail.next = appendNode
		s.tail = appendNode
	}
	return nil
}

func (s *SingleLinkedList[T]) InsertFirst(data T) error {
	node := &Node[T]{data, s.head}
	defer func() { s.len += 1 }()
	if s.len == 0 {
		s.tail = node
		s.head = node
	} else {
		s.head = node
	}
	return nil
}

func (s *SingleLinkedList[T]) InsertAfter(insertData T, data T) error {
	appendNode := &Node[T]{insertData, nil}
	if s.len == 0 {
		return errors.New("unable to insert after data when linked list is empty")
	}
	node := s.head
	for node != nil {
		if node.data == data {
			s.len += 1
			appendNode.next = node.next
			node.next = appendNode
			return nil
		}
		node = node.next
	}
	return errors.New("provided data not found in linked list... value is not inserted")
}
