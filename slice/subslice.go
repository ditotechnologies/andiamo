/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

func SubsliceAfter[Elem any](slice []Elem, idx int) []Elem {
	if idx < 0 {
		panic("cannot take slice after negative index")
	}
	if idx >= len(slice) {
		// no slice after
		return nil
	}
	output := make([]Elem, len(slice)-idx)
	for i := idx; i < len(slice); i += 1 {
		output[i-idx] = slice[i]
	}
	return output
}
