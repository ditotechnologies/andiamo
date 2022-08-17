/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package linkedlist

type Queue[T any] interface {
	Len() int
	Add(T)
	Remove() T
	Peek() T
}

func (ll *linkedList[T]) Add(elem T) {
	ll.AddFirst(elem)
}

func (ll *linkedList[T]) Remove() T {
	return ll.removeAtTail()
}

func (ll *linkedList[T]) Peek() T {
	return ll.valAtTail()
}

func NewQueue[T any]() Queue[T] {
	return newLinkedList[T]()
}
