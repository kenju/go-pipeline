// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package pipeline

import "context"

// GeneratorInterface generates channels from interface{} array
// Use ctx to cancel the stream processing.
func GeneratorInterface(
	ctx context.Context,
	values ...interface{},
) <-chan interface{} {
	ch := make(chan interface{}, len(values))

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

// GeneratorBool generates channels from bool array
// Use ctx to cancel the stream processing.
func GeneratorBool(
	ctx context.Context,
	values ...bool,
) <-chan bool {
	ch := make(chan bool, len(values))

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

// GeneratorByte generates channels from byte array
// Use ctx to cancel the stream processing.
func GeneratorByte(
	ctx context.Context,
	values ...byte,
) <-chan byte {
	ch := make(chan byte, len(values))

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

// GeneratorInt32 generates channels from int32 array
// Use ctx to cancel the stream processing.
func GeneratorInt32(
	ctx context.Context,
	values ...int32,
) <-chan int32 {
	ch := make(chan int32, len(values))

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

// GeneratorUint generates channels from uint array
// Use ctx to cancel the stream processing.
func GeneratorUint(
	ctx context.Context,
	values ...uint,
) <-chan uint {
	ch := make(chan uint, len(values))

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

// GeneratorUint64 generates channels from uint64 array
// Use ctx to cancel the stream processing.
func GeneratorUint64(
	ctx context.Context,
	values ...uint64,
) <-chan uint64 {
	ch := make(chan uint64, len(values))

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

// GeneratorFloat64 generates channels from float64 array
// Use ctx to cancel the stream processing.
func GeneratorFloat64(
	ctx context.Context,
	values ...float64,
) <-chan float64 {
	ch := make(chan float64, len(values))

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
