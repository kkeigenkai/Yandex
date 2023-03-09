package postfixnotation

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kkeigenkai/Yandex/util"
)

func PostfixNotation(str string) (answ int) {
	sr := strings.Split(str, " ")
	stack := util.CreateStack[string](len(sr))

	for _, c := range sr {
		if _, err := strconv.Atoi(c); err == nil {
			stack.Push(c)
		} else {
			xs, ok := stack.Pop()
			if !ok {
				panic("Stack is empty")
			}
			ys, ok := stack.Pop()
			if !ok {
				panic("Stack is empty")
			}
			x, err := strconv.Atoi(xs)
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(ys)
			if err != nil {
				panic(err)
			}
			switch c {
			case "+":
				stack.Push(fmt.Sprintf("%d", y+x))
			case "*":
				stack.Push(fmt.Sprintf("%d", y*x))
			case "/":
				stack.Push(fmt.Sprintf("%d", y/x))
			case "-":
				stack.Push(fmt.Sprintf("%d", y-x))
			}
		}
	}
	answStr, ok := stack.Pop()
	if !ok {
		panic("NaN")
	}
	answ, err := strconv.Atoi(answStr)
	if err != nil {
		panic(err)
	}
	return
}
