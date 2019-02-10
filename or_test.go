package pipeline_test

import (
	"fmt"
	"github.com/kenju/go-pipeline"
	"time"
)

func ExampleOrGenType() {
	sig := func(after time.Duration) <-chan pipeline.GenType {
		c := make(chan pipeline.GenType)
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	channels := []<-chan pipeline.GenType{
		sig(2 * time.Hour),
		sig(5 * time.Minute),
		sig(1 * time.Second), // the close message of this channel will close the whole channels
		sig(1 * time.Hour),
		sig(1 * time.Minute),
	}

	<-pipeline.OrGenType(channels...)
	fmt.Printf("done after %v", time.Since(start))
}
