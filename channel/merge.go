/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package channel

import (
	"context"
	"sync"
)

func MergeToSlice[T any](ctx context.Context, chs []<-chan T) []T {
	combined := Merge(ctx, chs)
	return ToSlice(combined)
}

func Merge[T any](ctx context.Context, chs []<-chan T) <-chan T {
	output := make(chan T)
	go func() {
		defer close(output)
		var wg sync.WaitGroup
		wg.Add(len(chs))
		for _, _ch := range chs {
			go func(ch <-chan T) {
				defer wg.Done()
				for elem := range ch {
					select {
					case output <- elem:
					case <-ctx.Done():
						return
					}
				}
			}(_ch)
		}
		wg.Wait()
	}()
	return output
}
