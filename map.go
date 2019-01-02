package pipeline

import "context"

// MapInt calculate fn(v int) to each values from values argument.
// Use ctx to cancel the stream processing.
func MapInt(
	ctx context.Context,
	fn func(v int) int,
	values <-chan int,
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
