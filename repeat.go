package main

import (
	"context"
)

func Repeat(
	ctx context.Context,
	values ...interface{},
) <-chan interface{} {
	valueCh := make(chan interface{})

	go func() {
		defer close(valueCh)

		for {
			for _, v := range values {
				select {
				case <-ctx.Done():
					return
				case valueCh <- v:
				}
			}
		}
	}()

	return valueCh
}

func RepeatFn(
	ctx context.Context,
	fn func() interface{},
) <-chan interface{} {
	valueCh := make(chan interface{})

	go func() {
		defer close(valueCh)

		for {
			select {
			case <-ctx.Done():
				return
			case valueCh <- fn():
			}
		}
	}()

	return valueCh
}
