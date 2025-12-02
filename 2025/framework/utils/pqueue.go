package utils

import "fmt"

type PriorityQueue[T any] struct {
	heap Heap[T]
}

func NewPriorityQueue[T any](comp Comparator[T]) PriorityQueue[T] {
	return PriorityQueue[T]{NewHeap[T](comp)}
}

func (q PriorityQueue[T]) Empty() bool {
	return q.heap.Empty()
}

func (q *PriorityQueue[T]) Enqueue(elm T) {
	q.heap.Push(elm)
}

func (q *PriorityQueue[T]) Dequeue() (T, error) {
	return q.heap.Pop()
}

func (q PriorityQueue[T]) Peek() (T, error) {
	return q.heap.Peek()
}

func TestPriorityQueue() int {
	q := NewPriorityQueue[int](func(a, b int) int { return a - b })

	if !q.Empty() {
		fmt.Println("New queue should be empty.")
		return 1
	}

	q.Enqueue(2)
	q.Enqueue(1)
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
		fmt.Println("PriorityQueue should be empty after dequeueing everything")
		return 1
	}

	elm, err = q.Dequeue()

	if err == nil {
		fmt.Println("Should not be able to dequeue from an empty queue.")
		return 1
	}
	return 0
}
