/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package channel

import (
	"context"
	"sync"
)

// WaitOrError Waits for a slice of errors to finish, or, returns the error if one of them errors. Only returns the first
// error to happen
func WaitOrError(_ctx context.Context, chs []<-chan error) error {

	output := make(chan error)

	ctx, cancel := context.WithCancel(_ctx)
	defer cancel()

	go func() {
		var wg sync.WaitGroup
		defer close(output)
		for _, _ch := range chs {
			wg.Add(1)
			go func(ch <-chan error) {
				defer wg.Done()
				for e := range ch {
					select {
					case output <- e:
					case <-ctx.Done():
					}
				}
			}(_ch)
		}
		wg.Wait()
	}()

	for elem := range output {
		if elem != nil {
			return elem
		}
	}

	return nil
}
