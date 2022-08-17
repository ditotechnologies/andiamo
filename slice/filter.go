/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

func filterElemToCh[Elem any](elem Elem, filterFn func(Elem) bool) <-chan Elem {
	ch := make(chan Elem)
	go func() {
		defer close(ch)
		if filterFn(elem) {
			ch <- elem
		}
	}()
	return ch
}

func Filter[Elem any](slice []Elem, filterFn func(Elem) bool) []Elem {
	chs := make([]<-chan Elem, len(slice))
	for idx, elem := range slice {
		chs[idx] = filterElemToCh(elem, filterFn)
	}
	output := make([]Elem, 0)
	for _, ch := range chs {
		for elem := range ch {
			output = append(output, elem)
		}
	}
	return output
}
