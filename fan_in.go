package pipeline

import (
	"sync"
	"context"
)

// FanIn multiplex multiple channels.
// Use ctx to cancel the stream processing.
func FanIn(
	ctx context.Context,
	channels ...<-chan interface{},
) <-chan interface{} {
	var wg sync.WaitGroup

	// select from all the channels
	wg.Add(len(channels))

	multiplexedCh := make(chan interface{})
	multiplex := func(c <-chan interface{}) {
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
