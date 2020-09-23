package llstack

import "errors"

type node struct {
	Data interface{}
	Next *node
}

// Stack represents the linked list implementation of the stack.
type Stack struct {
	length int
	Head   *node
}

// InitStack initializes and returns a reference to a new stack instance.
func InitStack() *Stack {
	return &Stack{}
}

// Size returns the size of the stack.
func (s *Stack) Size() int {
	return s.length
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

// Push adds a new element at the top of the stack.
func (s *Stack) Push(data interface{}) {
	node := &node{
		Data: data,
	}
	if s.Head == nil {
		s.Head = node
	} else {
		node.Next = s.Head
		s.Head = node
	}
	s.length++
}

// Pop deletes the top most element of the stack and returns it or returns an error if the stack is empty.
func (s *Stack) Pop() (interface{}, error) {
	if s.length == 0 {
		return nil, errors.New("unable to pop. Stack is empty")
	}
	data := s.Head.Data
	temp := s.Head.Next
	s.Head.Next = nil
	s.Head = temp
	s.length--
	return data, nil
}

// Peek returns the top most element of the stack or an error if the stack is empty.
func (s *Stack) Peek() (interface{}, error) {
	if s.length == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.Head.Data, nil
}
