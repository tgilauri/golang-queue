package queue

import (
	"fmt"
	"testing"
)

func TestQueueShouldHasPassedParams(t *testing.T) {
	size := 10
	queue := NewQueue[string](size, false)

	if queue.GetSize() != size {
		t.Errorf("Queue should has size of %d. Queue size is %d", size, queue.GetSize())
	}
}

func TestQueueShouldPrependElements(t *testing.T) {
	queue := NewQueue[string](10, false)
	amount := 5

	for i := 0; i < amount; i++ {
		queue.PushLeft(fmt.Sprintf("string%d", i))
	}

	if queue.IsEmpty() {
		t.Errorf("Queue should contain elements. Queue length is %d", queue.GetLength())
	}

	if queue.GetLength() != amount {
		t.Errorf("Queue should contain %d elements. Queue length is %d", amount, queue.GetLength())
	}
}

func TestQueueShouldAppendElements(t *testing.T) {
	queue := NewQueue[string](10, false)
	amount := 5

	for i := 0; i < amount; i++ {
		queue.PushRight(fmt.Sprintf("string-%d", i))
	}

	if queue.IsEmpty() {
		t.Errorf("Queue should contain elements. Queue length is %d", queue.GetLength())
	}

	if queue.GetLength() != amount {
		t.Errorf("Queue should contain %d elements. Queue length is %d", amount, queue.GetLength())
	}

	if queue.GetFirst() != "string-0" {
		t.Errorf("Elements should start with 0. First element is %s", queue.GetFirst())
	}
}

func TestQueueShouldBeEmptyAfterClean(t *testing.T) {
	queue := NewQueue[string](1, false)

	queue.PushRight("string")

	if queue.IsEmpty() {
		t.Errorf("Queue should contain elements. Queue length is %d", queue.GetLength())
	}

	queue.Clear()

	if !queue.IsEmpty() {
		t.Errorf("Queue should be empty. Queue length is %d", queue.GetLength())
	}
}

func TestQueueShouldBeFilledFromArray(t *testing.T) {
	arr := []string{"string-1", "string-2", "string-3"}
	queue := FromArray[string](arr, false)

	if queue.IsEmpty() {
		t.Errorf("Queue should contain elements. Queue length is %d", queue.GetLength())
	}

	if queue.GetLength() != len(arr) {
		t.Errorf("Queue should has length of source array. Queue length is %d", queue.GetLength())
	}
}
