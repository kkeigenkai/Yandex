package task18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kkeigenkai/yandex/util"
)

func Deque(cmds []string) []string {
	deque := util.CreateDeque[int]()
	res := make([]string, 0, len(cmds))

	for _, cmd := range cmds {
		switch cmdS := strings.Split(cmd, " "); cmdS[0] {
		case "push_front":
			x, err := strconv.Atoi(cmdS[1])
			if err != nil {
				panic(err)
			}
			deque.PushFront(x)
			res = append(res, "ok")

		case "push_back":
			x, err := strconv.Atoi(cmdS[1])
			if err != nil {
				panic(err)
			}
			deque.PushBack(x)
			res = append(res, "ok")

		case "pop_front":
			if deque.Size() == 0 {
				res = append(res, "error")
				continue
			}
			res = append(res, fmt.Sprintf("%d", deque.PopFront()))

		case "pop_back":
			if deque.Size() == 0 {
				res = append(res, "error")
				continue
			}
			res = append(res, fmt.Sprintf("%d", deque.PopBack()))

		case "front":
			if deque.Size() == 0 {
				res = append(res, "error")
				continue
			}
			res = append(res, fmt.Sprintf("%d", deque.Front()))

		case "back":
			if deque.Size() == 0 {
				res = append(res, "error")
				continue
			}
			res = append(res, fmt.Sprintf("%d", deque.Back()))

		case "size":
			res = append(res, fmt.Sprintf("%d", deque.Size()))

		case "clear":
			deque.Clear()
			res = append(res, "ok")

		case "exit":
			return append(res, "bye")

		}
	}

	return res
}
