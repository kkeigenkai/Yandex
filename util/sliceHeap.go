package util

type SliceHeap []int

func CreateSliceHeap() *SliceHeap {
	h := make(SliceHeap, 0)
	return &h
}

func (h *SliceHeap) Push(v int) {
	*h = append(*h, v)
	pos := len(*h) - 1
	for pos > 0 && (*h)[pos] > (*h)[(pos-1)/2] {
		(*h)[pos], (*h)[(pos-1)/2] = (*h)[(pos-1)/2], (*h)[pos]
		pos = (pos - 1) / 2
	}
}

func (h *SliceHeap) Pop() (v int) {
	v = (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	pos := 0

	for pos*2+2 < len(*h) {
		maxPos := leftSon(pos)
		if (*h)[maxPos] < (*h)[rightSon(pos)] {
			maxPos = rightSon(pos)
		}
		if (*h)[pos] < (*h)[maxPos] {
			(*h)[pos], (*h)[maxPos] = (*h)[maxPos], (*h)[pos]
			pos = maxPos
		} else {
			break
		}
	}

	*h = (*h)[:len(*h)-1]
	return v
}

func leftSon(i int) int {
	return i*2 + 1
}

func rightSon(i int) int {
	return i*2 + 2
}
