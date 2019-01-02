package pipeline

import (
	"context"
)

// Repeat return value via channel from values argument.
// Use ctx to cancel the stream processing.
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

// RepeatFn call fn() via channel.
// Use ctx to cancel the stream processing.
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
