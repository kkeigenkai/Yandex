package task11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kkeigenkai/yandex/util"
)

func StackErr(cmds []string) []string {
	res := make([]string, 0, len(cmds))
	stack := util.CreateLinkedStack[int]()
	for _, cmd := range cmds {
		cs := strings.Split(cmd, " ")
		if len(cs) == 2 {
			x, err := strconv.Atoi(cs[1])
			if err != nil {
				panic(err)
			}
			stack.Push(x)
			res = append(res, "ok")
			continue
		}
		switch cmd {
		case "pop":
			if stack.Cnt == 0 {
				res = append(res, "error")
				continue
			}
			res = append(res, fmt.Sprintf("%d", stack.Pop()))
		case "back":
			if stack.Cnt == 0 {
				res = append(res, "error")
				continue
			}
			res = append(res, fmt.Sprintf("%d", stack.Val))
		case "size":
			res = append(res, fmt.Sprintf("%d", stack.Cnt))
		case "clear":
			stack.Clear()
			res = append(res, "ok")
		case "exit":
			res = append(res, "bye")
			return res
		}
	}
	return res
}
