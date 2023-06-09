/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

func Reverse[Elem any](slice []Elem) {
	stopIdx := len(slice) / 2
	if len(slice)%2 == 1 {
		// odd we have to be minus 1
		stopIdx += 1
	}
	for fromIdx := 0; fromIdx < stopIdx; fromIdx += 1 {
		toIdx := (len(slice) - 1) - fromIdx
		slice[fromIdx], slice[toIdx] = slice[toIdx], slice[fromIdx]
	}
}
