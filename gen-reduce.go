// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package pipeline

import "context"

// ReduceInterface reduce values to the accumulator.
// Use ctx to cancel the stream processing.
func ReduceInterface(
	ctx context.Context,
	fn func(v, acc interface{}) interface{},
	values <-chan interface{},
) <-chan interface{} {
	ch := make(chan interface{})

	var acc interface{}
	go func() {
		defer close(ch)

		for v := range values {
			select {
			case <-ctx.Done():
				return
			default:
				acc = fn(v, acc)
				ch <- acc
			}
		}
	}()

	return ch
}

// ReduceString reduce values to the accumulator.
// Use ctx to cancel the stream processing.
func ReduceString(
	ctx context.Context,
	fn func(v, acc string) string,
	values <-chan string,
) <-chan string {
	ch := make(chan string)

	var acc string
	go func() {
		defer close(ch)

		for v := range values {
			select {
			case <-ctx.Done():
				return
			default:
				acc = fn(v, acc)
				ch <- acc
			}
		}
	}()

	return ch
}

// ReduceInt reduce values to the accumulator.
// Use ctx to cancel the stream processing.
func ReduceInt(
	ctx context.Context,
	fn func(v, acc int) int,
	values <-chan int,
) <-chan int {
	ch := make(chan int)

	var acc int
	go func() {
		defer close(ch)

		for v := range values {
			select {
			case <-ctx.Done():
				return
			default:
				acc = fn(v, acc)
				ch <- acc
			}
		}
	}()

	return ch
}

// ReduceFloat32 reduce values to the accumulator.
// Use ctx to cancel the stream processing.
func ReduceFloat32(
	ctx context.Context,
	fn func(v, acc float32) float32,
	values <-chan float32,
) <-chan float32 {
	ch := make(chan float32)

	var acc float32
	go func() {
		defer close(ch)

		for v := range values {
			select {
			case <-ctx.Done():
				return
			default:
				acc = fn(v, acc)
				ch <- acc
			}
		}
	}()

	return ch
}
