/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package linkedlist

type Deque[T any] interface {
	Len() int
	RemoveFirst() T
	PeekFirst() T
	AddFirst(T)
	RemoveLast() T
	AddLast(T)
	PeekLast() T
}

func (ll *linkedList[T]) AddFirst(val T) {
	ll.addAtHead(val)
}

func (ll *linkedList[T]) AddLast(val T) {
	ll.addAtTail(val)
}

func (ll *linkedList[T]) PeekFirst() T {
	return ll.valAtHead()
}

func (ll *linkedList[T]) PeekLast() T {
	return ll.valAtTail()
}

func (ll *linkedList[T]) RemoveFirst() T {
	return ll.removeAtHead()
}

func (ll *linkedList[T]) RemoveLast() T {
	return ll.removeAtTail()
}

func NewDeque[T any]() Deque[T] {
	return newLinkedList[T]()
}
