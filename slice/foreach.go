/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import "sync"

func ForEach[Elem any](slice []Elem, fn func(Elem)) {
	var wg sync.WaitGroup
	wg.Add(len(slice))
	for _, _elem := range slice {
		go func(elem Elem) {
			defer wg.Done()
			fn(elem)
		}(_elem)
	}
	wg.Wait()
}
