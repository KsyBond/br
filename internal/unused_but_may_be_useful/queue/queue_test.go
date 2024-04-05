package queue

import (
	"testing"
)

func Test_New(T *testing.T) {
	var q Queue[int]

	q.Enqueue(5)

	if q.Dequeue() != 5 {
		T.Fail()
	}
}
