package pipeline_test

import (
	"testing"
	"fmt"
	"context"
	"github.com/kenju/go-pipeline"
)

func TestMapInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mapFn := func(v int) int { return v + 1 }
	for v := range pipeline.MapInt(ctx, mapFn, pipeline.GeneratorInt(ctx, 1, 2, 3, 4)) {
		fmt.Println(v)
	}
}

func TestMapString(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mapFn := func(v string) string { return "**" + v + "**"}
	for v := range pipeline.MapString(ctx, mapFn, pipeline.GeneratorString(ctx, "hello", "world")) {
		fmt.Println(v)
	}
}
