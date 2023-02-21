package util

import (
	"fmt"
	"strings"
)

type LinkedStack[T any] struct {
	Val  T
	Cnt  int
	Prev *LinkedStack[T]
}

func CreateLinkedStack[T any]() *LinkedStack[T] {
	s := &LinkedStack[T]{}
	return s
}

func (s *LinkedStack[T]) String() string {
	head := s
	ss := make([]string, 0, head.Cnt)
	for head.Cnt != 0 && head != nil {
		ss = append(ss, fmt.Sprintf("%v", head.Val))
		head = head.Prev
	}
	return strings.Join(ss, " -> ")
}

func (s *LinkedStack[T]) IsEmpty() bool {
	return s.Cnt == 0
}

func (s *LinkedStack[T]) Push(v T) {
	p := *s
	s.Val = v
	s.Cnt = p.Cnt + 1
	s.Prev = &p
}

func (s *LinkedStack[T]) Pop() T {
	v := s.Val
	if s.Cnt == 0 {
		return v
	}
	s.Val = s.Prev.Val
	s.Cnt = s.Prev.Cnt
	s.Prev = s.Prev.Prev
	return v
}

func (s *LinkedStack[T]) GetLast() (v T) {
	return s.Val
}

func (s *LinkedStack[T]) Clear() {
	var v T
	s.Val = v
	s.Cnt = 0
	s.Prev = nil
}
