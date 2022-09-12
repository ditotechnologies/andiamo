/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package channel

import (
	"context"
	"github.com/ditotechnologies/andiamo/either"
	"sync"
)

// ParallelizeFunctions Runs functions in parallel and consolidates the result to a
// channel
func ParallelizeFunctions[T any](ctx context.Context, fns []func() T) <-chan T {
	output := make(chan T)
	go func() {
		defer close(output)
		var wg sync.WaitGroup
		for _, _fn := range fns {
			wg.Add(1)
			go func(fn func() T) {
				defer wg.Done()
				result := fn()
				select {
				case output <- result:
				case <-ctx.Done():
				}
			}(_fn)
		}
		wg.Wait()
	}()
	return output
}

// ParallelizeFunctionsToResultAndError Runs functions in parallel and then outputs the result to an array or an error.
// results are ordered
func ParallelizeFunctionsToResultAndError[T any](_ctx context.Context, fns []func() (T, error)) ([]T, error) {
	ctx, cancel := context.WithCancel(_ctx)
	defer cancel()

	channels := make([]chan either.Either[T, error], 0)

	for _, _fn := range fns {
		_ch := make(chan either.Either[T, error])
		go func(ch chan either.Either[T, error], fn func() (T, error)) {
			defer close(ch)
			select {
			case <-ctx.Done():
			case _ch <- either.FunctionResultOrError(fn):
			}
		}(_ch, _fn)
		channels = append(channels, _ch)
	}

	collected := make([]T, 0)
	for _, ch := range channels {
		result := <-ch
		if !result.IsLeft() {
			return nil, result.Right()
		}
		collected = append(collected, result.Left())
	}
	return collected, nil
}

func ParallelizeFunctionsToError(_ctx context.Context, fns []func() error) error {
	ctx, cancel := context.WithCancel(_ctx)
	defer cancel()
	ch := ParallelizeFunctions[error](ctx, fns)
	for err := range ch {
		if err != nil {
			return err
		}
	}
	return nil
}
