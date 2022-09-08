/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import "github.com/ditotechnologies/andiamo/set"

// Deduplicate Removes all duplicate elements of a slice. Returns items in the order they appear.
func Deduplicate[Elem comparable](slice []Elem) []Elem {
	return DeduplicateWithKey[Elem, Elem](slice, func(e Elem) Elem {
		return e
	})
}

// DeduplicateWithKey Deduplicate Removes all duplicate elements of a slice. Returns items in the order they appear. An element is a
// duplicate if they have the same key as defined by the key function
func DeduplicateWithKey[Elem any, Key comparable](slice []Elem, key func(Elem) Key) []Elem {
	output := make([]Elem, 0)
	exists := set.New[Key]()
	for _, elem := range slice {
		k := key(elem)
		if !exists.Contains(k) {
			output = append(output, elem)
			exists.Add(k)
		}
	}
	return output
}
