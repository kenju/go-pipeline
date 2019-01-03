package pipeline_test

import (
	"testing"
	"context"
	"github.com/kenju/go-pipeline"
	"fmt"
)

func TestReduceInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reduceFn := func(v, acc int) int { return v + acc }
	for v := range pipeline.ReduceInt(ctx, reduceFn, pipeline.GeneratorInt(ctx, 1, 2, 3, 4, 5)) {
		fmt.Println(v)
	}
}

func TestReduceString(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reduceFn := func(v, acc string) string { return acc + "!!" + v }
	for v := range pipeline.ReduceString(ctx, reduceFn, pipeline.GeneratorString(ctx, "hello", "world")) {
		fmt.Println(v)
	}
}

func TestReduceFloat32(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reduceFn := func(v, acc float32) float32 { return v + acc }
	for v := range pipeline.ReduceFloat32(ctx, reduceFn, pipeline.GeneratorFloat32(ctx, 1.1, 2.2, 3.3)) {
		fmt.Println(v)
	}
}
