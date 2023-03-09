package task19

import "github.com/kkeigenkai/yandex/util"

func Heap(ss [][2]int) []int {
	h := util.CreateSliceHeap()
	res := make([]int, 0, len(ss))
	for _, cmd := range ss {
		switch cmd[0] {
		case 0:
			h.Push(cmd[1])
		case 1:
			res = append(res, h.Pop())
		}
	}
	return res
}
