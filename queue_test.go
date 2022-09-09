package evaluator

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := newQueue()

	if q == nil {
		t.Fatal("queue should be created")
	}
}

func TestQueueEnqueue(t *testing.T) {
}

func TestQueueDequeue(t *testing.T) {
}
