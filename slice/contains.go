/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

func Contains[Elem comparable](slice []Elem, value Elem) bool {
	_, ok := IndexOf(slice, value)
	return ok
}
