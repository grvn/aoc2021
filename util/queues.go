package util

import "container/list"

type Queue struct {
	values *list.List
}

func NewQueue() *Queue {
	return &Queue{
		values: list.New(),
	}
}

func (queue *Queue) Length() int {
	return queue.values.Len()
}

func (queue *Queue) IsEmpty() bool {
	return queue.values.Len() == 0
}

// Returns value and true or
// nil and false indicating no value
func (queue *Queue) Peek() (interface{}, bool) {
	if queue.IsEmpty() {
		return nil, false
	} else {
		return queue.values.Back().Value, true
	}
}

func (queue *Queue) Push(item interface{}) {
	queue.values.PushBack(item)
}

// Returns value and true or
// nil and false indicating no value
func (queue *Queue) Pop() (interface{}, bool) {
	if queue.IsEmpty() {
		return nil, false
	} else {
		return queue.values.Remove(queue.values.Back()), true
	}
}
