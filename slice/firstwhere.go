/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

func FirstWhere[Elem any](slice []Elem, fn func(Elem) bool) (elem Elem, ok bool) {
	idx, hasIdx := IndexOfWithComparator(slice, fn)
	if !hasIdx {
		return
	}
	return slice[idx], true
}
