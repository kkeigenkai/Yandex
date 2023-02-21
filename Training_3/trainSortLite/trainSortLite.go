package trainsortlite

import (
	"github.com/kkeigenkai/Yandex/util"
)

func SortLite(trainCnt int, trains []int) bool {
	queue := util.CreateLinkedStack[int]()
	path2 := util.CreateLinkedStack[int]()
	path2.Push(0)
	for _, t := range trains {
		queue.Push(t)
		for queue.Cnt != 0 {
			if path2.Val+1 == queue.Val {
				path2.Push(queue.Val)
				queue.Pop()
			} else {
				break
			}
		}
	}
	return path2.Cnt == trainCnt+1
}
