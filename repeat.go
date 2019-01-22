package pipeline

import (
	"context"
)

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "GenType=interface{},byte,string,int,uint64,float32,float64"

// RepeatGenType return value via channel from values argument.
// Use ctx to cancel the stream processing.
func RepeatGenType(
	ctx context.Context,
	values ...GenType,
) <-chan GenType {
	ch := make(chan GenType)

	go func() {
		defer close(ch)

		for {
			for _, v := range values {
				select {
				case <-ctx.Done():
					return
				case ch <- v:
				}
			}
		}
	}()

	return ch
}

// RepeatFnGenType call fn() via channel.
// Use ctx to cancel the stream processing.
func RepeatFnGenType(
	ctx context.Context,
	fn func() GenType,
) <-chan GenType {
	ch := make(chan GenType)

	go func() {
		defer close(ch)

		for {
			select {
			case <-ctx.Done():
				return
			case ch <- fn():
			}
		}
	}()

	return ch
}
