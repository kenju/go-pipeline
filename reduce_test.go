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
