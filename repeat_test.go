package pipeline_test

import (
	"testing"
	"fmt"
	"context"
	"github.com/kenju/go-pipeline"
	"math/rand"
)

func TestRepeat(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range pipeline.Take(ctx, pipeline.Repeat(ctx, 1), 10) {
		fmt.Printf("%v ", v)
	}
}

func TestRepeatFn(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	randFn := func() interface{} { return rand.Int() }
	for v := range pipeline.Take(ctx, pipeline.RepeatFn(ctx, randFn), 5) {
		fmt.Printf("%v ", v)
	}
}

func TestRepeatString(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range pipeline.TakeString(ctx, pipeline.RepeatString(ctx, "hello"), 10) {
		fmt.Printf("%v ", v)
	}
}

func TestRepeatFnString(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repeatFn := func() string { return "!" }
	for v := range pipeline.TakeString(ctx, pipeline.RepeatFnString(ctx, repeatFn), 5) {
		fmt.Printf("%v ", v)
	}
}

func TestRepeatInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range pipeline.TakeInt(ctx, pipeline.RepeatInt(ctx, 1), 10) {
		fmt.Printf("%v ", v)
	}
}

func TestRepeatFnInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repeatFn := func() int { return 9 }
	for v := range pipeline.TakeInt(ctx, pipeline.RepeatFnInt(ctx, repeatFn), 5) {
		fmt.Printf("%v ", v)
	}
}

func TestRepeatFloat32(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range pipeline.TakeFloat32(ctx, pipeline.RepeatFloat32(ctx, 1.1), 10) {
		fmt.Printf("%v ", v)
	}
}

func TestRepeatFnFloat32(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repeatFn := func() float32 { return 1.1 }
	for v := range pipeline.TakeFloat32(ctx, pipeline.RepeatFnFloat32(ctx, repeatFn), 5) {
		fmt.Printf("%v ", v)
	}
}
