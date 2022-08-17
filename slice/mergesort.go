/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

func mergeSortStep[Elem any](slice []Elem, comparator func(Elem, Elem) bool, startIdx int, endIdx int) <-chan Elem {
	output := make(chan Elem)
	go func() {
		defer close(output)
		if endIdx == startIdx {
			// stopping condition, just one element
			output <- slice[endIdx]
			return
		}
		midPoint := (endIdx + startIdx) / 2
		c1 := mergeSortStep(slice, comparator, startIdx, midPoint)
		c2 := mergeSortStep(slice, comparator, midPoint+1, endIdx)

		var c1Val Elem
		c1HasVal := false
		var c2Val Elem
		c2HasVal := false

		for {
			if !c1HasVal {
				c1Read, c1Ok := <-c1
				if c1Ok {
					c1Val = c1Read
					c1HasVal = true
				}
			}
			if !c2HasVal {
				c2Read, c2Ok := <-c2
				if c2Ok {
					c2Val = c2Read
					c2HasVal = true
				}
			}
			if !c1HasVal && !c2HasVal {
				// stopping condition, nothing more to sort.
				break
			} else if !c2HasVal {
				// just need to send c1
				output <- c1Val
				c1HasVal = false
			} else if !c1HasVal {
				// just need to send c2
				output <- c2Val
				c2HasVal = false
			} else if comparator(c1Val, c2Val) {
				// comparator says use c1
				output <- c1Val
				c1HasVal = false
			} else {
				// comparator says use c2
				output <- c2Val
				c2HasVal = false
			}
		}

	}()

	return output
}

func MergeSort[Elem any](slice []Elem, comparator func(Elem, Elem) bool) []Elem {
	if len(slice) <= 1 {
		return slice
	}
	ch := mergeSortStep(slice, comparator, 0, len(slice)-1)
	collected := make([]Elem, len(slice))
	idx := 0
	for elem := range ch {
		collected[idx] = elem
		idx += 1
	}
	return collected
}
