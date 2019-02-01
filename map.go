package pipeline

import (
	"context"
)

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "GenType=interface{},bool,byte,string,int,uint64,float32,float64"

// MapGenType calculate fn(v GenType) to each values from values argument.
// Use ctx to cancel the stream processing.
func MapGenType(
	ctx context.Context,
	fn func(v GenType) GenType,
	values <-chan GenType,
) <-chan GenType {
	ch := make(chan GenType)

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
