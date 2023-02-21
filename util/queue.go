package util

import (
	"fmt"
	"strings"
)

type LinkedQueue[T any] struct {
	First *QueueItem[T]
	Last  *QueueItem[T]
	Size  int
}

type QueueItem[T any] struct {
	Val  T
	Next *QueueItem[T]
}

func CreateLinkedQueue[T any](st []T) *LinkedQueue[T] {
	q := &LinkedQueue[T]{}
	if len(st) != 0 {
		for _, v := range st {
			q.Add(v)
		}
	}
	return q
}

func createQueueItem[T any](v T) *QueueItem[T] {
	return &QueueItem[T]{
		Val: v,
	}
}

func (q LinkedQueue[T]) String() string {
	head := q.First
	res := make([]string, 0, q.Size)
	for head != nil {
		res = append(res, fmt.Sprintf("%v", head.Val))
		head = head.Next
	}
	return strings.Join(res, " <-- ")
}

func (q *LinkedQueue[T]) Add(v T) {
	qi := createQueueItem(v)
	if q.Size == 0 {
		q.First = qi
		q.First.Next = qi
		q.Last = qi
	}
	q.Last.Next = qi
	q.Last = qi
	q.Size++
}

func (q *LinkedQueue[T]) Remove() (v T) {
	if q.Size == 0 {
		return
	}
	v = q.First.Val
	q.First = q.First.Next
	q.Size--
	return
}

func (q *LinkedQueue[T]) Front() (v T) {
	if q.Size != 0 {
		v = q.First.Val
	}
	return
}

func (q *LinkedQueue[T]) Back() (v T) {
	if q.Size != 0 {
		v = q.Last.Val
	}
	return
}
