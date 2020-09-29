package arrqueue

import "errors"

type Queue []interface{}

func CreateQueue() Queue {
	return Queue{}
}

func (q *Queue) Enqueue(data interface{}) {
	*q = append((*q), data)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Peek() interface{} {
	return (*q)[0]
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("pop failed. Queue is empty")
	}
	index := len(*q) - 1
	data := (*q)[0]
	*q = (*q)[:index]
	return data, nil
}
