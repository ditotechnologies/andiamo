/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package linkedlist

import "sync"

type linkedList[T any] struct {
	accessLock sync.RWMutex
	head       *node[T]
	tail       *node[T]
	len        int
}

type node[T any] struct {
	next     *node[T]
	previous *node[T]
	val      T
}

func (ll *linkedList[T]) Len() int {
	ll.accessLock.RLock()
	defer ll.accessLock.RUnlock()
	return ll.len
}

func newLinkedList[T any]() *linkedList[T] {
	return &linkedList[T]{
		len: 0,
	}
}

func (ll *linkedList[T]) removeAtHead() T {
	ll.accessLock.Lock()
	defer ll.accessLock.Unlock()
	if ll.len == 0 {
		panic("cannot remove first of an empty linked list")
	}
	val := ll.head.val
	if ll.len == 1 {
		ll.head = nil
		ll.tail = nil
	} else {
		ll.head.next.previous = nil
		ll.head = ll.head.next
	}
	ll.len -= 1
	return val
}

func (ll *linkedList[T]) removeAtTail() T {
	ll.accessLock.Lock()
	defer ll.accessLock.Unlock()
	if ll.len == 0 {
		panic("canot remove first of an empty linked list")
	}
	val := ll.head.val
	if ll.len == 1 {
		ll.head = nil
		ll.tail = nil
	} else {
		ll.tail.previous.next = nil
		ll.tail = ll.tail.previous
	}
	ll.len -= 1
	return val
}

func (ll *linkedList[T]) valAtHead() T {
	ll.accessLock.RLock()
	defer ll.accessLock.RUnlock()
	if ll.len == 0 {
		panic("cannot see the head value of an empty linked list")
	}
	return ll.head.val
}

func (ll *linkedList[T]) valAtTail() T {
	ll.accessLock.RLock()
	defer ll.accessLock.RUnlock()
	if ll.len == 0 {
		panic("cannot see the tail value of an empty linked list")
	}
	return ll.tail.val
}

func (ll *linkedList[T]) addAtHead(val T) {
	ll.accessLock.Lock()
	defer ll.accessLock.Unlock()
	ll.len += 1
	if ll.head == nil {
		ll.head = &node[T]{
			val: val,
		}
		ll.tail = ll.head
	} else {
		newNode := &node[T]{
			next: ll.head,
			val:  val,
		}
		ll.head.previous = newNode
		ll.head = newNode
	}
}

func (ll *linkedList[T]) addAtTail(val T) {
	ll.accessLock.Lock()
	defer ll.accessLock.Unlock()
	ll.len += 1
	if ll.tail == nil {
		ll.head = &node[T]{
			val: val,
		}
		ll.tail = ll.head
	} else {
		newNode := &node[T]{
			previous: ll.tail,
			val:      val,
		}
		ll.tail.next = newNode
		ll.tail = newNode
	}
}
