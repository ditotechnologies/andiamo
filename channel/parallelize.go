/*
 * Copyright (c) Dito Technologies LLC. All rights reserved.
 */

package channel

import (
	"context"
	"github.com/ditotechnologies/andiamo/either"
	"github.com/ditotechnologies/andiamo/slice"
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
// results are not ordered
func ParallelizeFunctionsToResultAndError[T any](_ctx context.Context, fns []func() (T, error)) ([]T, error) {
	ctx, cancel := context.WithCancel(_ctx)
	defer cancel()
	wrappedFunctions := slice.Map[func() (T, error), func() either.Either[T, error]](fns, func(fn func() (T, error)) func() either.Either[T, error] {
		return func() either.Either[T, error] {
			return either.FunctionResultOrError(fn)
		}
	})
	ch := ParallelizeFunctions[either.Either[T, error]](ctx, wrappedFunctions)
	collected := make([]T, 0)
	for elem := range ch {
		if !elem.IsLeft() {
			return nil, elem.Right()
		}
		collected = append(collected, elem.Left())
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
