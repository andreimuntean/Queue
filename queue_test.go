package queue_test

import (
	"queue"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	assert := assert.New(t)

	_, firstErr := queue.New(0);
	_, secondErr := queue.New(1);
	_, thirdErr := queue.New(5);
	_, fourthErr := queue.New(1000000);

	assert.NotNil(firstErr, "Does not accept non-positive capacity.")
	assert.Nil(secondErr, "Can create queue with 1 element.")
	assert.Nil(thirdErr, "Can create queue with 5 elements.")
	assert.Nil(fourthErr, "Can create queue with 1000000 elements.")
}

func TestCap(t *testing.T) {
	assert := assert.New(t)

	firstQueue, _ := queue.New(1);
	secondQueue, _ := queue.New(5);
	thirdQueue, _ := queue.New(10);
	fourthQueue, _ := queue.New(100);

	assert.Equal(1, firstQueue.Cap(), "Returns the correct capacity.")
	assert.Equal(5, secondQueue.Cap(), "Returns the correct capacity.")
	assert.Equal(10, thirdQueue.Cap(), "Returns the correct capacity.")
	assert.Equal(100, fourthQueue.Cap(), "Returns the correct capacity.")
}

func TestLen(t *testing.T) {
	assert := assert.New(t)

	queue, _ := queue.New(5);
	assert.Equal(queue.Len(), 0, "Returns the correct length.")

	queue.Push(0)
	assert.Equal(queue.Len(), 1, "Returns the correct length.")

	queue.Push(0, 0, 0, 0)
	assert.Equal(queue.Len(), 5, "Returns the correct length.")

	queue.Push(0, 0, 0)
	queue.Pop()
	assert.Equal(queue.Len(), 4, "Returns the correct length.")

	queue.Pop()
	queue.Pop()
	queue.Pop()
	queue.Pop()
	queue.Push(0)
	assert.Equal(queue.Len(), 1, "Returns the correct length.")

	queue.Pop()
	assert.Equal(queue.Len(), 0, "Returns the correct length.")

	queue.Pop()
	assert.Equal(queue.Len(), 0, "Returns the correct length.")
}

func TestPush(t *testing.T) {
	assert := assert.New(t)

	queue, _ := queue.New(5)

	queue.Push(10)
	assert.Equal(1, queue.Len(), "Push enqueues a value.")
	assert.Equal(10, queue.Peek(), "Push enqueues a value.")

	queue.Push(20)
	assert.Equal(2, queue.Len(), "Push enqueues a value.")
	assert.Equal(10, queue.Peek(), "Push enqueues a value.")
	
	queue.Push(30, 40, 50)
	assert.Equal(5, queue.Len(), "Push enqueues multiple values.")
	
	queue.Push(99)
	assert.Equal(5, queue.Len(), "Push no longer adds values when the queue is full.")

	for _, expectedValue := range []int { 10, 20, 30, 40, 50 } {
		actualValue := queue.Pop()
		assert.Equal(expectedValue, actualValue, "Push enqueues a value.")
	}
}