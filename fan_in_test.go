package pipeline_test

import (
	"context"
	"fmt"
	"github.com/kenju/go-pipeline"
)

func ExampleFanInInterface() {
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

func ExampleFanInString() {
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

func ExampleFanInInt() {
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

func ExampleFanInFloat32() {
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
