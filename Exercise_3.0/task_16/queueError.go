package task16

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kkeigenkai/yandex/util"
)

func QueueError(cmds []string) []string {
	queue := util.CreateLinkedQueue([]int{})
	res := make([]string, 0, len(cmds))
	for _, cmd := range cmds {
		switch cmdS := strings.Split(cmd, " "); cmdS[0] {
		case "pop":
			if queue.Size == 0 {
				res = append(res, "error")
				continue
			}
			res = append(res, fmt.Sprintf("%d", queue.Remove()))

		case "front":
			if queue.Size == 0 {
				res = append(res, "error")
				continue
			}
			res = append(res, fmt.Sprintf("%d", queue.Front()))

		case "size":
			res = append(res, fmt.Sprintf("%d", queue.Size))

		case "clear":
			queue.Clear()
			res = append(res, "ok")

		case "push":
			x, err := strconv.Atoi(cmdS[1])
			if err != nil {
				panic(err)
			}
			queue.Add(x)
			res = append(res, "ok")

		case "exit":
			return append(res, "bye")
		}
	}
	return res
}
