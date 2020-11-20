package elementary

import (
	"errors"
)

// Various errors a heap function can return.
var (
	ErrHeapOverflow  = errors.New("heap overflow")
	ErrHeapUnderflow = errors.New("heap underflow")
	ErrKeyMismatch   = errors.New("key is smaller than current")
)

func NewMaxHeap() *MaxHeap {
	h := make([]int, 0)
	return &MaxHeap{
		heap: h,
		size: 0,
	}
}

// MaxHeap defines a heap structure with a int slice and a heap sizes.
type MaxHeap struct {
	heap []int
	size int
}

// Parent returns the index of the parent of an element at index i.
func (h *MaxHeap) Parent(i int) int {
	return (i - 1) >> 1
}

// Left returns the index of the left child of an element at index i.
// If none, -1 and an error is returned.
func (h *MaxHeap) Left(i int) (int, error) {
	idx := (i << 1) + 1
	if idx > h.size-1 {
		return -1, ErrHeapOverflow
	}

	return idx, nil
}

// Right returns the index of the right child of an element at index i.
// If none, -1 and an error is returned.
func (h *MaxHeap) Right(i int) (int, error) {
	idx := (i << 1) + 2
	if idx > h.size-1 {
		return -1, ErrHeapOverflow
	}

	return idx, nil
}

// BuildHeap builds up the heap with a given slice of element values while
// ensuring the max heap property.
func (h *MaxHeap) BuildHeap(vals []int) {
	h.heap = vals
	h.size = len(vals)

	for i := h.size / 2; i >= 0; i-- {
		h.Heapify(i)
	}
}

// Heapify ensures the max heap property is maintained from a given index in the
// heap.
func (h *MaxHeap) Heapify(i int) error {
	if i > h.size-1 {
		return ErrHeapOverflow
	}

	lgst := i

	l, err := h.Left(i)
	if err == nil && h.heap[l] > h.heap[lgst] {
		lgst = l
	}

	r, err := h.Right(i)
	if err == nil && h.heap[r] > h.heap[lgst] {
		lgst = r
	}

	if lgst != i {
		h.heap[i], h.heap[lgst] = h.heap[lgst], h.heap[i]
		h.Heapify(lgst)
	}

	return nil
}

// Max returns the maximum element value of the heap.
func (h *MaxHeap) Max() int {
	return h.heap[0]
}

// ExtractMax removes and returns the maximum element value of the heap.
func (h *MaxHeap) ExtractMax() (int, error) {
	if h.size < 1 {
		return -1, ErrHeapUnderflow
	}

	max := h.heap[0]
	h.size--
	h.heap[0] = h.heap[h.size]
	h.Heapify(0)

	return max, nil
}

// IncreaseVal increases the element value for a given index.
func (h *MaxHeap) IncreaseVal(i, val int) error {
	if i > h.size-1 {
		return ErrHeapOverflow
	}

	if val < h.heap[i] {
		return ErrKeyMismatch
	}

	h.heap[i] = val
	for i > 0 {
		p := h.Parent(i)
		if h.heap[p] < h.heap[i] {
			h.heap[i], h.heap[p] = h.heap[p], h.heap[i]
		}
		i = p
	}

	return nil
}

// Insert adds an element value to the heap, and ensures the max heap property
// is still satisfied.
func (h *MaxHeap) Insert(val int) {
	h.heap = append(h.heap, -1)
	h.size++
	h.IncreaseVal(h.size-1, val)
}
