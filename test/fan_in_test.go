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
		channels[i] = pipeline.RepeatInterface(ctx, i)
	}

	for v := range pipeline.TakeInterface(ctx, pipeline.FanInInterface(ctx, channels...), 10) {
		fmt.Printf("%v ", v)
	}
}

func TestFanInString(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numChannels := 3
	channels := make([]<-chan string, numChannels)
	for i := 0; i < numChannels; i++ {
		channels[i] = pipeline.RepeatString(ctx, fmt.Sprintf("id:%d", i))
	}

	for v := range pipeline.TakeString(ctx, pipeline.FanInString(ctx, channels...), 10) {
		fmt.Printf("%v ", v)
	}
}

func TestFanInInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numChannels := 3
	channels := make([]<-chan int, numChannels)
	for i := 0; i < numChannels; i++ {
		channels[i] = pipeline.RepeatInt(ctx, i)
	}

	for v := range pipeline.TakeInt(ctx, pipeline.FanInInt(ctx, channels...), 10) {
		fmt.Printf("%v ", v)
	}
}

func TestFanInFloat32(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numChannels := 3
	channels := make([]<-chan float32, numChannels)
	for i := 0; i < numChannels; i++ {
		channels[i] = pipeline.RepeatFloat32(ctx, float32(i + 1) * 0.1)
	}

	for v := range pipeline.TakeFloat32(ctx, pipeline.FanInFloat32(ctx, channels...), 10) {
		fmt.Printf("%v ", v)
	}
}
