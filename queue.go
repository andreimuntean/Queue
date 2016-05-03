package queue

import (
	"errors"
)

type Queue struct {
	values []interface{}
}

func New(capacity int) (*Queue, error) {
	if capacity < 1 {
		return nil, errors.New("Queue capacity must be a positive integer.")
	}

	queue := new(Queue)
	queue.values = make([]interface{}, capacity)

	return queue, nil
}