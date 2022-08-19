/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package linkedlist

type Deque[T any] interface {
	// Len The number of items in the Deque
	Len() int
	// RemoveFirst Removes the first item, returning the value
	RemoveFirst() T
	// PeekFirst Returns the first item
	PeekFirst() T
	// AddFirst Adds an item to the start of the deque
	AddFirst(T)
	// RemoveLast Removes the last item, returning the value
	RemoveLast() T
	// AddLast Adds an item to the end of the deque
	AddLast(T)
	// PeekLast Removes the last item from the deque
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
