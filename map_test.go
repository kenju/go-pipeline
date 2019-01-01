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
