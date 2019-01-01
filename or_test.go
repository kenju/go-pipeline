package pipeline_test

import (
	"testing"
	"time"
	"fmt"
	"github.com/kenju/go-pipeline"
)

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	channels := []<-chan interface{}{
		sig(2 * time.Hour),
		sig(5 * time.Minute),
		sig(1 * time.Second), // the close message of this channel will close the whole channels
		sig(1 * time.Hour),
		sig(1 * time.Minute),
	}

	<-pipeline.Or(channels...)
	fmt.Printf("done after %v", time.Since(start))
}
