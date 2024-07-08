package main

import (
	"fmt"
	"math"
)

/*
To maintain always correct binary tree, we adjust depth to 2^depth+depth.

Example:
Depth := 0    (0)

Depth := 1    (0)

		     /   \
	       (0)   (0)

Depth := 2            (0)

	               /       \
			     (0)       (0)
			    /   \     /   \
			  (0)   (0) (0)   (0)

Etc.
*/
func heapDepthFormula(depth float64) float64 {
	if depth == 0 {
		return depth
	}
	return math.Pow(2, depth) + heapDepthFormula(depth-1)
}

// child L 2i+1
func leftChildIdx(parentIdx int) int {
	return 2*parentIdx + 1
}

// child R 2i+2
func rightChildIdx(parentIdx int) int {
	return 2*parentIdx + 2
}

// parent L/R (i-1)/2
func parentIdx(childIdx int) int {
	return (childIdx - 1) / 2
}

type minHeap struct {
	h []int
}

func NewMinHeap(depth float64) *minHeap {
	if depth < 0 {
		depth = heapDepthFormula(2)
	} else {
		depth = heapDepthFormula(depth)
	}
	return &minHeap{make([]int, 0, int(depth))}
}

func (mh *minHeap) Insert(value int) {
	mh.h = append(mh.h, value)
	mh.heapifyUp()
}
func (mh *minHeap) Delete(valueAt int) (deleted int) {
	if len(mh.h) == 0 {
		return -1
	}
	if len(mh.h) == 1 {
		deleted = mh.h[0]
		mh.h = mh.h[:0]
		return
	}
	deleted = mh.h[valueAt]
	mh.h[valueAt], mh.h[len(mh.h)-1] = mh.h[len(mh.h)-1], mh.h[valueAt]
	mh.h = mh.h[:len(mh.h)-1]
	mh.heapifyDown(valueAt)
	return
}

func (mh *minHeap) heapifyUp() {
	lastInsertedIdx := len(mh.h) - 1
	heapifyTo := parentIdx(lastInsertedIdx)
	for heapifyTo >= 0 {
		if mh.h[heapifyTo] <= mh.h[lastInsertedIdx] {
			break
		}
		mh.h[heapifyTo], mh.h[lastInsertedIdx] = mh.h[lastInsertedIdx], mh.h[heapifyTo]
		lastInsertedIdx = heapifyTo
		heapifyTo = parentIdx(lastInsertedIdx)
	}
}

func (mh *minHeap) heapifyDown(valueAt int) {
	for {
		leftIdx := leftChildIdx(valueAt)
		rightIdx := rightChildIdx(valueAt)
		if leftIdx >= len(mh.h) || rightIdx >= len(mh.h) {
			break
		}
		if mh.h[leftIdx] < mh.h[rightIdx] && mh.h[valueAt] > mh.h[leftIdx] {
			mh.h[valueAt], mh.h[leftIdx] = mh.h[leftIdx], mh.h[valueAt]
			valueAt = leftIdx
			continue
		}
		if mh.h[valueAt] > mh.h[rightIdx] {
			mh.h[valueAt], mh.h[rightIdx] = mh.h[rightIdx], mh.h[valueAt]
			valueAt = rightIdx
			continue
		}
		break
	}
}

func main() {
	mh := NewMinHeap(2)
	mh.Insert(10)
	mh.Insert(5)
	mh.Insert(2)
	mh.Insert(1000)
	mh.Insert(2000)
	mh.Insert(990)
	mh.Insert(900)
	mh.Insert(500)
	mh.Delete(1)
	mh.Delete(3)
	for i, v := range mh.h {
		fmt.Printf("ind %d val %d\n", i, v)

	}
}
