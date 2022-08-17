/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

func IndexOf[Elem comparable](slice []Elem, value Elem) (int, bool) {
	for idx, elem := range slice {
		if elem == value {
			return idx, true
		}
	}
	return -1, false
}

func IndexOfWithComparator[Elem any](slice []Elem, fn func(Elem) bool) (int, bool) {
	for idx, elem := range slice {
		if fn(elem) {
			return idx, true
		}
	}
	return -1, false
}
