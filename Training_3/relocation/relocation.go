package relocation

import (
	"github.com/kkeigenkai/Yandex/util"
)

func Relocation(cnt int, si []int) []int {
	stack := util.CreateLinkedStack[[2]int]()
	for i := 1; i < len(si); i++ {
		if si[i] >= stack.Val[0] {
			stack.Push([2]int{si[i], i})
		} else {
			for stack.Cnt != 0 && stack.Val[0] > si[i] {
				v := stack.Pop()
				si[v[1]] = i
			}
			stack.Push([2]int{si[i], i})
		}
	}
	for stack.Cnt != 0 {
		v := stack.Pop()
		si[v[1]] = -1
	}
	return si
}
