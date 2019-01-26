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

func ExampleFanInByte() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numChannels := 3
	channels := make([]<-chan byte, numChannels)
	for i := 0; i < numChannels; i++ {
		channels[i] = pipeline.RepeatByte(ctx, 0x01, 0x02, 0x03)
	}

	for v := range pipeline.TakeByte(ctx, pipeline.FanInByte(ctx, channels...), 10) {
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

func ExampleFanInUint64() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numChannels := 3
	channels := make([]<-chan uint64, numChannels)
	for i := 0; i < numChannels; i++ {
		channels[i] = pipeline.RepeatUint64(ctx, uint64(i))
	}

	for v := range pipeline.TakeUint64(ctx, pipeline.FanInUint64(ctx, channels...), 10) {
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

func ExampleFanInFloat64() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numChannels := 3
	channels := make([]<-chan float64, numChannels)
	for i := 0; i < numChannels; i++ {
		channels[i] = pipeline.RepeatFloat64(ctx, float64(i + 1) * 0.1)
	}

	for v := range pipeline.TakeFloat64(ctx, pipeline.FanInFloat64(ctx, channels...), 10) {
		fmt.Printf("%v ", v)
	}
}
