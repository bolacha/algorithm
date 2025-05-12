package minheap

import "fmt"

type MinHeap struct {
	array []int
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MinHeap) Swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *MinHeap) Insert(i int) {
	h.array = append(h.array, i)

	h.HeapifyUp(len(h.array) - 1)
}

func (h *MinHeap) HeapifyUp(index int) {
	for h.array[parent(index)] > h.array[index] {
		h.Swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) Remove() int {
	removed := h.array[0]

	l := len(h.array) - 1

	if len(h.array) == 0 {
		return -1
	}

	h.array[0] = h.array[l]

	h.array = h.array[:l]

	h.HeapifyDown(0)

	return removed
}

func (h *MinHeap) HeapifyDown(index int) {
	lastIndexToCheck := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndexToCheck {
		if l == lastIndexToCheck {
			childToCompare = l
		} else if h.array[l] < h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] > h.array[childToCompare] {
			h.Swap(index, childToCompare)

			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MinHeap) Print() {
	fmt.Println(h.array)
}
