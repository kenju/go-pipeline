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
		fmt.Println(v)
	}
}
