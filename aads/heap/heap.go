package main

import "fmt"

type heap struct {
	arr []int
}

func (h *heap) siftUp(i int) {
	for i > 0 {
		p := (i - 1) / 2          // parent
		if h.arr[p] <= h.arr[i] { // if parent less than child break
			break
		}
		h.arr[p], h.arr[i] = h.arr[i], h.arr[p] // else swap parent with child
		i = p                                   // update indexes
	}
}

func (h *heap) siftDown(i int) {
	n := len(h.arr)
	for {
		l := 2*i + 1 // left child
		if l >= n {
			break
		}
		j := l                            // left child
		r := l + 1                        // right child
		if r < n && h.arr[r] < h.arr[l] { // if right child less than left
			// (less than parent) then
			j = r // we need to swap with right child
		}
		if h.arr[i] <= h.arr[j] { // if parent less than left/right child then stop
			break
		}
		h.arr[i], h.arr[j] = h.arr[j], h.arr[i] // swap elements
		i = j                                   // update indexes
	}
}

func (h *heap) push(x int) {
	h.arr = append(h.arr, x) // add new element
	h.siftUp(len(h.arr) - 1) // sift up new element before see smaller parent
}

func (h *heap) pop() (int, bool) {
	n := len(h.arr) // get count of elements
	if n == 0 {     // is zero so heap is empty
		return 0, false
	}
	res := h.arr[0]     // result as minimum element
	last := h.arr[n-1]  // last element
	h.arr = h.arr[:n-1] // cut last element [0; n-1)
	if n-1 > 0 {        // if count without min greater than 0
		// so can
		h.arr[0] = last // set last as top element
		h.siftDown(0)   // sift down top element before see greater child
	}
	return res, true
}

func main() {
	heap := &heap{arr: make([]int, 0)}
	heap.push(1)
	heap.push(2)
	heap.push(3)
	heap.push(-2)
	heap.push(-1)
	heap.push(0)
	heap.push(-3)
	v, isV := heap.pop()
	for isV {
		fmt.Println(v)
		v, isV = heap.pop()
	}
}
