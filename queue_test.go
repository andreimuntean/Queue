package queue_test

import (
	"queue"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	assert := assert.New(t)

	_, firstErr := queue.New(0);
	assert.NotNil(firstErr, "Does not accept non-positive capacity.")

	_, secondErr := queue.New(1);
	assert.Nil(secondErr, "Can create queue with 1 element.")

	_, thirdErr := queue.New(5);
	assert.Nil(thirdErr, "Can create queue with 5 elements.")

	_, fourthErr := queue.New(1000000);
	assert.Nil(fourthErr, "Can create queue with 1000000 elements.")
}