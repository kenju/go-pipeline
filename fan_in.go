package pipeline

import (
	"sync"
	"context"
)

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "GenType=interface{},byte,string,int,uint64,float32,float64"

// FanInGenType multiplex multiple channels.
// Use ctx to cancel the stream processing.
func FanInGenType(
	ctx context.Context,
	channels ...<-chan GenType,
) <-chan GenType {
	var wg sync.WaitGroup

	// select from all the channels
	wg.Add(len(channels))

	multiplexedCh := make(chan GenType)
	multiplex := func(c <-chan GenType) {
		defer wg.Done()
		for i := range c {
			select {
			case <-ctx.Done():
				return
			case multiplexedCh <- i:
			}
		}
	}
	for _, c := range channels {
		go multiplex(c)
	}

	// close a channel for multiplexing when all channels are processed
	go func() {
		wg.Wait()
		close(multiplexedCh)
	}()

	return multiplexedCh
}
