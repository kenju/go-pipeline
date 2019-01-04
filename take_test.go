package pipeline_test

import (
	"testing"
	"github.com/kenju/go-pipeline"
	"fmt"
	"context"
)

func TestTake(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range pipeline.Take(ctx, pipeline.Repeat(ctx, 1), 3) {
		fmt.Printf("%v ", v)
	}
}

func TestTakeInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range pipeline.TakeInt(ctx, pipeline.GeneratorInt(ctx, 1, 2, 3), 1) {
		fmt.Printf("%v ", v)
	}
}

func TestTakeString(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range pipeline.TakeString(ctx, pipeline.GeneratorString(ctx, "hello", "world", "!"), 1) {
		fmt.Printf("%v ", v)
	}
}

func TestTakeFloat32(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range pipeline.TakeFloat32(ctx, pipeline.GeneratorFloat32(ctx, 1.1, 2.2, 3.3), 1) {
		fmt.Printf("%v ", v)
	}
}
