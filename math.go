package pipeline

import "context"

// Add calculate addition to each values from intCh.
// Use ctx to cancel the stream processing.
func Add(
	ctx context.Context,
	intCh <-chan int,
	additive int,
) <-chan int {
	addedCh := make(chan int)

	go func() {
		defer close(addedCh)

		for i := range intCh {
			select {
			case <-ctx.Done():
				return
			case addedCh <- i + additive:
			}
		}
	}()

	return addedCh
}

// Multiply calculate multiply to each values from intCh.
// Use ctx to cancel the stream processing.
func Multiply(
	ctx context.Context,
	intCh <-chan int,
	multiplier int,
) <-chan int {
	multipliedCh := make(chan int)

	go func() {
		defer close(multipliedCh)

		for i := range intCh {
			select {
			case <-ctx.Done():
				return
			case multipliedCh <- i * multiplier:
			}
		}
	}()

	return multipliedCh
}
