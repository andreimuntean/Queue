package queue

import (
	"errors"
)

type Queue struct {
	values []interface{}
	nextPopIndex int
	nextPushIndex int
	length int
}

func New(capacity int) (*Queue, error) {
	if capacity < 1 {
		return nil, errors.New("Queue capacity must be a positive integer.")
	}

	queue := new(Queue)
	queue.values = make([]interface{}, capacity)

	return queue, nil
}

func (queue *Queue) Cap() int {
	return cap(queue.values)
}

func (queue *Queue) Len() int {
	return queue.length	
}

func (queue *Queue) IsEmpty() bool {
	return queue.Len() == 0
}

func (queue *Queue) IsFull() bool {
	return queue.Len() == queue.Cap()
}

func (queue *Queue) Push(values ...interface{}) {
	for _, value := range values {
		if queue.IsFull() {
			// Does not accept new values.
			return
		}

		// Adds the value to the queue.
		queue.values[queue.nextPushIndex] = value
		queue.nextPushIndex = (queue.nextPushIndex + 1) % len(queue.values)
		queue.length++
	}
}

func (queue *Queue) Peek() interface{} {
	if queue.IsEmpty() {
		return nil
	}

	return queue.values[queue.nextPopIndex]
}

func (queue *Queue) Pop() interface{} {
	if queue.IsEmpty() {
		return nil
	}

	// Removes the value from the queue.
	value := queue.values[queue.nextPopIndex]
	queue.nextPopIndex = (queue.nextPopIndex + 1) % len(queue.values)
	queue.length--

	return value
}