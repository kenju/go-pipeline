package pipeline

import (
	"context"
)

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "GenType=interface{},bool,byte,string,int,int32,uint,uint64,float32,float64"

// TakeGenType return n of GenType items from valueCh channel.
// Use ctx to cancel the stream processing.
func TakeGenType(
	ctx context.Context,
	valueCh <-chan GenType,
	num int,
) <-chan GenType {
	ch := make(chan GenType)

	go func() {
		defer close(ch)

		for i := 0; i < num; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- <-valueCh:
			}
		}
	}()

	return ch
}
