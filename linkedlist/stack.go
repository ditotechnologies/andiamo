/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package linkedlist

type Stack[T any] interface {
	Len() int
	Pop() T
	Push(T)
	Peek() T
}

// note: peek defined for queue as valAtTail

func (ll *linkedList[T]) Pop() T {
	return ll.removeAtTail()
}

func (ll *linkedList[T]) Push(val T) {
	ll.addAtTail(val)
}

func NewStack[T any]() Stack[T] {
	return newLinkedList[T]()
}
