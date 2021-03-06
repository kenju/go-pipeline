// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package pipeline

import "context"

// MapInterface calculate fn(v Interface) to each values from values argument.
// Use ctx to cancel the stream processing.
func MapInterface(
	ctx context.Context,
	fn func(v interface{}) interface{},
	values <-chan interface{},
) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)

		for v := range values {
			select {
			case <-ctx.Done():
				return
			case ch <- fn(v):
			}
		}
	}()

	return ch
}

// MapBool calculate fn(v Bool) to each values from values argument.
// Use ctx to cancel the stream processing.
func MapBool(
	ctx context.Context,
	fn func(v bool) bool,
	values <-chan bool,
) <-chan bool {
	ch := make(chan bool)

	go func() {
		defer close(ch)

		for v := range values {
			select {
			case <-ctx.Done():
				return
			case ch <- fn(v):
			}
		}
	}()

	return ch
}

// MapByte calculate fn(v Byte) to each values from values argument.
// Use ctx to cancel the stream processing.
func MapByte(
	ctx context.Context,
	fn func(v byte) byte,
	values <-chan byte,
) <-chan byte {
	ch := make(chan byte)

	go func() {
		defer close(ch)

		for v := range values {
			select {
			case <-ctx.Done():
				return
			case ch <- fn(v):
			}
		}
	}()

	return ch
}

// MapString calculate fn(v String) to each values from values argument.
// Use ctx to cancel the stream processing.
func MapString(
	ctx context.Context,
	fn func(v string) string,
	values <-chan string,
) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		for v := range values {
			select {
			case <-ctx.Done():
				return
			case ch <- fn(v):
			}
		}
	}()

	return ch
}

// MapUint calculate fn(v Uint) to each values from values argument.
// Use ctx to cancel the stream processing.
func MapUint(
	ctx context.Context,
	fn func(v uint) uint,
	values <-chan uint,
) <-chan uint {
	ch := make(chan uint)

	go func() {
		defer close(ch)

		for v := range values {
			select {
			case <-ctx.Done():
				return
			case ch <- fn(v):
			}
		}
	}()

	return ch
}

// MapInt calculate fn(v Int) to each values from values argument.
// Use ctx to cancel the stream processing.
func MapInt(
	ctx context.Context,
	fn func(v int) int,
	values <-chan int,
) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for v := range values {
			select {
			case <-ctx.Done():
				return
			case ch <- fn(v):
			}
		}
	}()

	return ch
}

// MapUint64 calculate fn(v Uint64) to each values from values argument.
// Use ctx to cancel the stream processing.
func MapUint64(
	ctx context.Context,
	fn func(v uint64) uint64,
	values <-chan uint64,
) <-chan uint64 {
	ch := make(chan uint64)

	go func() {
		defer close(ch)

		for v := range values {
			select {
			case <-ctx.Done():
				return
			case ch <- fn(v):
			}
		}
	}()

	return ch
}

// MapFloat32 calculate fn(v Float32) to each values from values argument.
// Use ctx to cancel the stream processing.
func MapFloat32(
	ctx context.Context,
	fn func(v float32) float32,
	values <-chan float32,
) <-chan float32 {
	ch := make(chan float32)

	go func() {
		defer close(ch)

		for v := range values {
			select {
			case <-ctx.Done():
				return
			case ch <- fn(v):
			}
		}
	}()

	return ch
}

// MapFloat64 calculate fn(v Float64) to each values from values argument.
// Use ctx to cancel the stream processing.
func MapFloat64(
	ctx context.Context,
	fn func(v float64) float64,
	values <-chan float64,
) <-chan float64 {
	ch := make(chan float64)

	go func() {
		defer close(ch)

		for v := range values {
			select {
			case <-ctx.Done():
				return
			case ch <- fn(v):
			}
		}
	}()

	return ch
}
