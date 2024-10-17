/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package slice

import (
	"context"
	"sync"
)

// Runs a function in parallel over a slice. That function can return an error. This resolved it to the one error.
func ForEachWithError[Elem any](slice []Elem, fn func(Elem) error) error {

	errCh := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		defer close(errCh)
		var wg sync.WaitGroup
		for _, _elem := range slice {
			wg.Add(1)
			go func(elem Elem) {
				defer wg.Done()
				result := fn(elem)
				select {
				case <-ctx.Done():
				case errCh <- result:
				}
			}(_elem)
		}
		wg.Wait()
	}()

	for elem := range errCh {
		if elem != nil {
			return elem
		}
	}

	return nil
}
