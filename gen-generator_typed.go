// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package pipeline

import "context"

// GeneratorString generates channels from string array
// Use ctx to cancel the stream processing.
func GeneratorString(
	ctx context.Context,
	values ...string,
) <-chan string {
	ch := make(chan string, len(values))

	go func() {
		defer close(ch)

		for _, v := range values {
			select {
			case <-ctx.Done():
				return
			case ch <- v:
			}
		}
	}()

	return ch
}

// GeneratorInt generates channels from int array
// Use ctx to cancel the stream processing.
func GeneratorInt(
	ctx context.Context,
	values ...int,
) <-chan int {
	ch := make(chan int, len(values))

	go func() {
		defer close(ch)

		for _, v := range values {
			select {
			case <-ctx.Done():
				return
			case ch <- v:
			}
		}
	}()

	return ch
}

// GeneratorFloat32 generates channels from float32 array
// Use ctx to cancel the stream processing.
func GeneratorFloat32(
	ctx context.Context,
	values ...float32,
) <-chan float32 {
	ch := make(chan float32, len(values))

	go func() {
		defer close(ch)

		for _, v := range values {
			select {
			case <-ctx.Done():
				return
			case ch <- v:
			}
		}
	}()

	return ch
}