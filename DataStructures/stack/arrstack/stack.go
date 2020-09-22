package arrstack

import (
	"errors"
)

type Stack []interface{}

func InitStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Push(data interface{}) {
	(*s) = append((*s), data)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("pop failed. Stack is empty")
	}
	index := len(*s) - 1
	data := (*s)[index]
	(*s) = (*s)[:index]
	return data, nil
}

func (s *Stack) Peek() (data interface{}) {
	data = (*s)[len(*s)-1]
	return data
}
