package util

import (
	"fmt"
	"strings"
)

type deque[T any] struct {
	front *dequeItem[T]
	back  *dequeItem[T]
	size  int
}

type dequeItem[T any] struct {
	val  T
	next *dequeItem[T]
	prev *dequeItem[T]
}

func CreateDeque[T any]() *deque[T] {
	d := &deque[T]{}
	return d
}

func (d *deque[T]) String() string {
	ds := make([]string, 0, d.size)
	h := d.front
	for i := 0; i < d.size; i++ {
		ds = append(ds, fmt.Sprintf("%v", h.val))
		h = h.prev
	}
	return strings.Join(ds, " <--> ")
}

func (d *deque[T]) Front() T {
	return d.front.val
}

func (d *deque[T]) Back() T {
	return d.back.val
}

func (d *deque[T]) Size() int {
	return d.size
}

func (d *deque[T]) PushFront(v T) {
	qi := &dequeItem[T]{
		val:  v,
		prev: d.front,
	}
	if d.size == 0 {
		d.front = qi
		d.back = qi
	} else {
		d.front.next = qi
		d.front = qi
	}
	d.size++
}

func (d *deque[T]) PushBack(v T) {
	qi := &dequeItem[T]{
		val:  v,
		next: d.back,
	}
	if d.size == 0 {
		d.front = qi
		d.back = qi
	} else {
		d.back.prev = qi
		d.back = qi
	}
	d.size++
}

func (d *deque[T]) PopFront() (v T) {
	switch d.size {
	case 0:
		return
	case 1:
		v = d.front.val
		d.front = nil
		d.back = nil
	default:
		v = d.front.val
		nf := d.front.prev
		d.front.prev = nil
		d.front = nf
		d.front.next = nil
	}
	d.size--
	return
}

func (d *deque[T]) PopBack() (v T) {
	switch d.size {
	case 0:
		return
	case 1:
		v = d.back.val
		d.front = nil
		d.back = nil
	default:
		v = d.back.val
		nb := d.back.next
		d.back.next = nil
		d.back = nb
		d.back.prev = nil

	}
	d.size--
	return
}

func (d *deque[T]) Clear() {
	d.front = nil
	d.back = nil
	d.size = 0
}
