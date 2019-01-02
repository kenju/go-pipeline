package pipeline

import (
	"context"
	)

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "GenType=string,int"

// GeneratorGenType generates channels from GenType array
// Use ctx to cancel the stream processing.
func GeneratorGenType(
	ctx context.Context,
	values ...GenType,
) <-chan GenType {
	ch := make(chan GenType, len(values))

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
