package utils

import "fmt"

type Queue[T any] struct {
	q []T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{make([]T, 0)}
}

func (q Queue[T]) Empty() bool {
	return len(q.q) == 0
}

func (q *Queue[T]) Enqueue(elm T) {
	q.q = append(q.q, elm)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.Empty() {
		var x T
		return x, fmt.Errorf("queue is empty")
	}

	elm := q.q[0]
	q.q = q.q[1:]
	return elm, nil
}

func (q Queue[T]) Peek() (T, error) {
	if q.Empty() {
		var x T
		return x, fmt.Errorf("queue is empty")
	}
	return q.q[0], nil
}

func TestQueue() int {
	q := NewQueue[int]()

	if !q.Empty() {
		fmt.Println("New queue should be empty.")
		return 1
	}

	q.Enqueue(1)
	q.Enqueue(2)
	elm, err := q.Peek()
	if err != nil {
		fmt.Printf("Failed to peek: %v\n", err)
		return 1
	}

	if elm != 1 {
		fmt.Printf("Peek returned the wrong value: %d\n", elm)
		return 1
	}

	elm, err = q.Dequeue()
	if err != nil {
		fmt.Printf("Failed to dequeue: %v\n", err)
		return 1
	}

	if elm != 1 {
		fmt.Printf("Dequeue returned the wrong value: %d\n", elm)
		return 1
	}

	elm, err = q.Peek()
	if err != nil {
		fmt.Printf("Failed to peek: %v\n", err)
		return 1
	}

	if elm != 2 {
		fmt.Printf("Peek returned the wrong value: %d\n", elm)
		return 1
	}

	elm, err = q.Dequeue()
	if err != nil {
		fmt.Printf("Failed to dequeue: %v\n", err)
		return 1
	}

	if elm != 2 {
		fmt.Printf("Dequeue returned the wrong value: %d\n", elm)
		return 1
	}

	if !q.Empty() {
		fmt.Println("Queue should be empty after dequeueing everything")
		return 1
	}

	elm, err = q.Dequeue()

	if err == nil {
		fmt.Println("Should not be able to dequeue from an empty queue.")
		return 1
	}
	return 0
}
