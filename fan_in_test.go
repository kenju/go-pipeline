package pipeline_test

import (
	"testing"
	"fmt"
	"context"
	"github.com/kenju/go-pipeline"
)

func TestFanIn(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numChannels := 3
	channels := make([]<-chan interface{}, numChannels)
	for i := 0; i < numChannels; i++ {
		channels[i] = pipeline.Repeat(ctx, i)
	}

	for v := range pipeline.Take(ctx, pipeline.FanIn(ctx, channels...), 10) {
		fmt.Println(v)
	}
}
