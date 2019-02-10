package pipeline_test

import (
	"context"
	"fmt"
	"github.com/kenju/go-pipeline"
)

func ExampleFanInGenType() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numChannels := 3
	channels := make([]<-chan pipeline.GenType, numChannels)
	for i := 0; i < numChannels; i++ {
		channels[i] = pipeline.RepeatGenType(ctx, i)
	}

	for v := range pipeline.TakeGenType(ctx, pipeline.FanInGenType(ctx, channels...), 10) {
		fmt.Printf("%v ", v)
	}
}
