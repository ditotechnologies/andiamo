/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import "github.com/ditotechnologies/andiamo/set"

// Deduplicate Removes all duplicate elements of a slice. Returns items in the order they appear.
func Deduplicate[Elem comparable](slice []Elem) []Elem {
	output := make([]Elem, 0)
	exists := set.New[Elem]()
	for _, elem := range slice {
		if !exists.Contains(elem) {
			output = append(output, elem)
			exists.Add(elem)
		}
	}
	return output
}
