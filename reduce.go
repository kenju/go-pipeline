package pipeline

import "context"

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
