package utils

// Reference: https://en.wikipedia.org/wiki/Binary_heap

import (
	"fmt"
	"strings"
)

type Heap[T any] struct {
	list []T
	comp Comparator[T]
}

func NewHeap[T any](comp Comparator[T]) Heap[T] {
	return Heap[T]{[]T{}, comp}
}

func (heap *Heap[T]) Push(values ...T) {
	if len(values) == 1 {
		heap.list = append(heap.list, values[0])
		heap.bubbleUp()
	} else {
		heap.list = append(heap.list, values...)

		size := len(heap.list)/2 + 1
		for i := size; i >= 0; i-- {
			heap.bubbleDownIndex(i)
		}
	}
}

func (heap *Heap[T]) bubbleUp() {
	idx := len(heap.list) - 1
	for idx > 0 {
		parentIdx := (idx - 1) / 2
		idxValue := heap.list[idx]
		parentValue := heap.list[parentIdx]
		if heap.comp(parentValue, idxValue) <= 0 {
			break
		}
		heap.list[idx], heap.list[parentIdx] = heap.list[parentIdx], heap.list[idx]
		idx = parentIdx
	}
}

func (heap *Heap[T]) bubbleDown() {
	heap.bubbleDownIndex(0)
}

func (heap *Heap[T]) bubbleDownIndex(idx int) {
	size := len(heap.list)
	leftIdx := idx*2 + 1
	for leftIdx < size {
		rightIdx := idx*2 + 2
		smallerIdx := leftIdx
		smallerValue := heap.list[leftIdx]
		if rightIdx < size {
			rightValue := heap.list[rightIdx]
			if heap.comp(smallerValue, rightValue) > 0 {
				smallerIdx = rightIdx
				smallerValue = rightValue
			}
		}
		idxValue := heap.list[idx]
		if heap.comp(idxValue, smallerValue) > 0 {
			heap.list[idx], heap.list[smallerIdx] = heap.list[smallerIdx], heap.list[idx]
		} else {
			break
		}
		idx = smallerIdx
		leftIdx = idx*2 + 1
	}
}

func (heap Heap[T]) Size() int {
	return len(heap.list)
}

func (heap Heap[T]) Empty() bool {
	return len(heap.list) == 0
}

func (heap Heap[T]) Peek() (T, error) {
	if heap.Empty() {
		var x T
		return x, fmt.Errorf("heap is empty")
	}
	return heap.list[0], nil
}

func (heap *Heap[T]) Pop() (T, error) {
	if heap.Empty() {
		var x T
		return x, fmt.Errorf("heap is empty")
	}

	val := heap.list[0]
	lastIdx := len(heap.list) - 1
	heap.list[0], heap.list[lastIdx] = heap.list[lastIdx], heap.list[0]
	heap.list = heap.list[:lastIdx]
	heap.bubbleDown()
	return val, nil
}

func (heap Heap[T]) String() string {
	str := "BinaryHeap: ["
	values := []string{}
	for i := 0; i < len(heap.list); i++ {
		values = append(values, fmt.Sprintf("%v", heap.list[i]))
	}
	str += strings.Join(values, ", ")
	str += "]"
	return str
}

func TestHeap() int {
	minHeap := NewHeap[int](func(a, b int) int { return a - b })
	if !minHeap.Empty() {
		fmt.Println("Empty heap should be empty")
		return 1
	}
	minHeap.Push(9)
	minHeap.Push(8)
	minHeap.Push(7)
	minHeap.Push(1, 2, 3, 7, 8, 8)
	if minHeap.Empty() {
		fmt.Println("Empty heap")
		return 1
	}

	if minHeap.Size() != 9 {
		fmt.Println("Heap doesn't have 9 elements")
		fmt.Println(minHeap)
		return 1
	}

	if mn, err := minHeap.Peek(); err != nil || mn != 1 {
		fmt.Println("Heap.Peak does not return minimum.")
		fmt.Println(minHeap)
		return 1
	}

	mn, err := minHeap.Pop()
	if err != nil || mn != 1 {
		fmt.Println("Heap.Pop does not return minimum.")
		fmt.Println(minHeap)
		return 1
	}

	for !minHeap.Empty() {
		nmn, err := minHeap.Pop()
		if err != nil || nmn < mn {
			fmt.Println("Heap.Pop does not return minimum. (loop)")
			fmt.Println(minHeap)
			return 1
		}
		mn = nmn
	}
	maxHeap := NewHeap[int](func(a, b int) int { return b - a })
	if !maxHeap.Empty() {
		fmt.Println("Empty heap should be empty")
		return 1
	}
	maxHeap.Push(9)
	maxHeap.Push(8)
	maxHeap.Push(7)
	maxHeap.Push(1, 2, 3, 7, 8, 8)
	if maxHeap.Empty() {
		fmt.Println("Empty heap")
		return 1
	}

	if maxHeap.Size() != 9 {
		fmt.Println("Heap doesn't have 9 elements")
		fmt.Println(maxHeap)
		return 1
	}

	if mx, err := maxHeap.Peek(); err != nil || mx != 9 {
		fmt.Println("Heap.Peak does not return maximum (maxHeap).")
		fmt.Println(maxHeap)
		return 1
	}

	mx, err := maxHeap.Pop()
	if err != nil || mx != 9 {
		fmt.Println("Heap.Pop does not return maximum (maxHeap).")
		fmt.Println(maxHeap)
		return 1
	}

	for !maxHeap.Empty() {
		nmx, err := maxHeap.Pop()
		if err != nil || nmx > mx {
			fmt.Println("Heap.Pop does not return maximum (maxHeap, loop).")
			fmt.Println(maxHeap)
			return 1
		}
		mx = nmx
	}
	return 0
}
