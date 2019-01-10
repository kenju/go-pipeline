package pipeline

import "context"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "GenType=interface{},string,int,float32"

// ReduceGenType reduce values to the accumulator.
// Use ctx to cancel the stream processing.
func ReduceGenType(
	ctx context.Context,
	fn func(v, acc GenType) GenType,
	values <-chan GenType,
) <-chan GenType {
	ch := make(chan GenType)

	var acc GenType
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
