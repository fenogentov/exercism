// Package linkedlist implements tools for working with doubly linked list.
package linkedlist

import (
	"errors"
)

// Node - list item.
type Node struct {
	Val  interface{}
	prev *Node
	next *Node
}

// List - linked list.
type List struct {
	head *Node
	tail *Node
	len  int
}

var ErrEmptyList error = errors.New("error list is empty")

// Next provides the following list item.
func (e *Node) Next() *Node {
	if e == nil {
		return nil
	}
	return e.next
}

// Prev provides the previous list item.
func (e *Node) Prev() *Node {
	if e == nil {
		return nil
	}
	return e.prev
}

// NewList makes a new list
func NewList(args ...interface{}) *List {
	l := new(List)
	for _, v := range args {
		l.PushBack(v)
	}
	return l
}

// PushFront adds an element to the beginning of the list.
func (l *List) PushFront(v interface{}) {
	newElem := Node{Val: v, prev: nil, next: l.head}

	if l.head != nil {
		l.head.prev = &newElem
	} else {
		l.tail = &newElem
	}
	l.head = &newElem
	l.len++
}

// PushBack adds an item to the end of the list.
func (l *List) PushBack(v interface{}) {
	newElem := Node{Val: v, prev: l.tail, next: nil}

	if l.tail != nil {
		l.tail.next = &newElem
	} else {
		l.head = &newElem
	}
	l.tail = &newElem
	l.len++
}

// PopFront removes the element at the beginning of the list.
func (l *List) PopFront() (interface{}, error) {
	if l.head == nil {
		return 0, ErrEmptyList
	}
	v := l.head.Val
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	l.len--
	return v, nil
}

// PopBack removes an element to the end of the list.
func (l *List) PopBack() (interface{}, error) {
	if l.tail == nil {
		return 0, ErrEmptyList
	}
	v := l.tail.Val
	l.tail = l.tail.prev
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	l.len--
	return v, nil
}

// Reverse swaps elements.
func (l *List) Reverse() *List {
	c := l.head
	for c != nil {
		c.prev, c.next = c.next, c.prev
		c = c.prev
	}
	l.head, l.tail = l.tail, l.head
	return l
}

// First gives the first element of the list.
func (l *List) First() *Node {
	return l.head
}

// Last gives the last element of the list.
func (l *List) Last() *Node {
	return l.tail
}
